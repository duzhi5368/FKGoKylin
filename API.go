package FKGoKylin

import (
	"fmt"
	"strconv"
)

// 扩展API请参考：http://kylin.apache.org/docs/howto/howto_use_restapi.html

// 登陆
func (k *FKKylin) Login() (code int, body []byte, err error) {
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/user/authentication")
	request.Header.SetMethod("POST")
	return k.do(request)
}

// 查询
func (k *FKKylin) QueryByStruct(tableName string, fields []string, where interface{}, offset int, limit int, isDebug bool) (*QueryResult, error) {
	querySQL, err := k.buildUpSQL(tableName, fields, where)
	if err != nil{
		return nil, err
	}
	if isDebug{
		fmt.Println(querySQL)
	}
	result, err := k.query(querySQL, offset, limit, isDebug)
	return result, err
}
func (k *FKKylin) QueryBySQL(sql string, offset int, limit int, isDebug bool) (*QueryResult, error) {
	result, err := k.query(sql, offset, limit, isDebug)
	return result, err
}

// 查询Tables列表
func (k *FKKylin) ListTables() (code int, body []byte, err error) {
	if k.ProjectName == ""{
		return 0, nil, fmt.Errorf("Project name shouldn't be empty.")
	}
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/tables_and_columns?project=" + k.ProjectName)
	request.Header.SetMethod("GET")
	return k.do(request)
}

// 查询Cube列表
func (k *FKKylin) ListCubes(offset int, limit int) (code int, body []byte, err error) {
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/cubes?projectName=" + k.ProjectName +
		"&offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit))
	request.Header.SetMethod("GET")
	return k.do(request)
}

// 查询一个Cube
func (k *FKKylin) GetCube(cubeName string) (code int, body []byte, err error) {
	if cubeName == "" {
		return 0, nil, fmt.Errorf("cube name shouldn't be empty")
	}
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/cubes/" + cubeName)
	request.Header.SetMethod("GET")
	return k.do(request)
}

// 获取一个Cube描述
func (k *FKKylin)GetCubeDesc(cubeName string) (code int, body []byte, err error) {
	if cubeName == "" {
		return 0, nil, fmt.Errorf("cube name shouldn't be empty")
	}
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/cube_desc/" + cubeName)
	request.Header.SetMethod("GET")
	return k.do(request)
}

// 获取一个Model
func (k *FKKylin) GetModel(modelName string) (code int, body []byte, err error) {
	if modelName == "" {
		return 0, nil, fmt.Errorf("model name shouldn't be empty")
	}
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/model/" + modelName)
	request.Header.SetMethod("GET")
	return k.do(request)
}