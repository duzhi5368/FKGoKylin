package FKGoKylin

import (
	"encoding/base64"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"
	"reflect"
)

type FKKylin struct{
	ProjectName string
	Url string
	UserName string
	Password string
}

func CreateFKKylin(	ProjectName string, Url string, UserName string, Password string) *FKKylin{
	p := &FKKylin{
		ProjectName: ProjectName,
		Url: Url,
		UserName: UserName,
		Password: Password,
	}
	return p
}

func (k *FKKylin) query(querySQL string, offset int, limit int, isDebug bool)(*QueryResult, error){

	queryCondition := &QueryCondition{
		SQL: querySQL,
		Limit: limit,
		Offset: offset,
		Project: k.ProjectName,
		AcceptPartial: false,
	}
	httpCode, responseBody, err := k.sendQuery(queryCondition)
	if err != nil{
		return nil, err
	}
	if httpCode != 200{
		return nil, fmt.Errorf("Kylin server return error code: " + strconv.Itoa(httpCode))
	}
	return k.parseQueryResult(responseBody)
}

func (k* FKKylin) sendQuery(query *QueryCondition)(code int, body []byte, err error){
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/query")
	request.Header.SetMethod("POST")
	request.Header.Set("Content-Type", "application/json")
	request.SetBody(query.ToBytes())
	return k.do(request)
}

func (k *FKKylin) buildUpSQL(tableName string, fields []string, where interface{}) (string, error){
	tableName = strings.ToUpper(tableName)
	querySQL := "select "
	if fields == nil{
		querySQL += " * "
	} else {
		for _, field := range fields{
			querySQL += strings.ToUpper(field) + ","
		}
		querySQL = querySQL[0 : len(querySQL)-1]
	}
	querySQL += " from " + tableName + " where 1=1 "

	vList := reflect.ValueOf(where).Elem()
	tList := reflect.TypeOf(where).Elem()
	for i := 0; i < vList.NumField(); i++{
		v := vList.Field(i)
		t := tList.Field(i)

		if t.Tag.Get("kylin") == "not_query_condition"{
			continue
		}
		if v.Type().String() == "int"{
			if t.Tag.Get("kylin") == "necessary_query_condition" && v.Interface().(int) == 0{
				return "", fmt.Errorf(t.Name + " shouldn't be empty")
			}
			if v.Interface().(int) != 0{
				querySQL = querySQL + " and " + tableName + "." + strings.ToUpper(t.Tag.Get("json")) +
					"=" + strconv.Itoa(v.Interface().(int))
			}
		} else if v.Type().String() == "int64"{
			if t.Tag.Get("kylin") == "necessary_query_condition" && v.Interface().(int64) == 0{
				return "", fmt.Errorf(t.Name + " shouldn't be empty")
			}
			if v.Interface().(int64) != 0{
				querySQL = querySQL + " and " + tableName + "." + strings.ToUpper(t.Tag.Get("json")) +
					"=" + strconv.Itoa((int)(v.Interface().(int64)))
			}
		} else if v.Type().String() == "string"{
			if t.Tag.Get("kylin") == "necessary_query_condition" && v.Interface().(string) == ""{
				return "", fmt.Errorf(t.Name + " shouldn't be empty")
			}
			if v.Interface().(string) != ""{
				querySQL = querySQL + " and " + tableName + "." + strings.ToUpper(t.Tag.Get("json")) +
					"='" + v.Interface().(string) + "'"
			}
		} else if v.Type().String() == "float64"{
			if t.Tag.Get("kylin") == "necessary_query_condition" && v.Interface().(float64) == 0.0{
				return "", fmt.Errorf(t.Name + " shouldn't be empty")
			}
			if v.Interface().(float64) != 0.0{
				querySQL = querySQL + " and " + tableName + "." + strings.ToUpper(t.Tag.Get("json")) +
					"=" + strconv.FormatFloat(v.Interface().(float64), 'g', 2, 64)
			}
		} else if v.Type().String() == "time_pair"{
			if t.Tag.Get("kylin") == "necessary_query_condition" && v.Interface().(TimePair).IsZero(){
				return "", fmt.Errorf(t.Name + " shouldn't be empty")
			}
			if !v.Interface().(TimePair).IsZero(){
				querySQL = querySQL + " and " + tableName + "." + strings.ToUpper(t.Tag.Get("json")) + ">='" +
					v.Interface().(TimePair).StartTime.Format("2006-01-02") + "'"
				querySQL = querySQL + " and " + tableName + "." + strings.ToUpper(t.Tag.Get("json")) + "<'" +
					v.Interface().(TimePair).EndTime.Format("2006-01-02") + "'"
			}
		}
	}
	return querySQL, nil
}

func (k *FKKylin) parseQueryResult(body []byte) (*QueryResult, error){
	if body == nil{
		return nil, fmt.Errorf("Empty http response body.")
	}
	qr := &QueryResult{}
	err := json.Unmarshal(body, qr)
	if err != nil{
		return nil, err
	}
	if qr.IsException{
		return nil, fmt.Errorf(qr.ExceptionMessage)
	}
	return qr, nil
}

func (k *FKKylin) calcAuth() string{
	auth := k.UserName + ":" + k.Password
	return "Basic "+ base64.StdEncoding.EncodeToString([]byte(auth))
}

func (k *FKKylin) genBaseRequest() *fasthttp.Request{
	request := fasthttp.AcquireRequest()
	request.Header.Set("Authorization", k.calcAuth())
	request.Header.Set("Content-Type", "application/json")
	return request
}

func (k *FKKylin) do(request *fasthttp.Request) (code int, body []byte, err error) {
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	err = fasthttp.Do(request, response)
	if err == nil {
		code = response.StatusCode()
		body = response.Body()
	}
	return
}