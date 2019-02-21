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
func (k *FKKylin) Query(tableName string, fields []string, where interface{}, offset int, limit int, isDebug bool) (result *QueryResult, err error) {
	kr, err := k.query(tableName, fields, where, offset, limit, isDebug)
	if err != nil{
		return nil, err
	}
	var metas []string
	for _, m := range kr.ColumnMetas{
		metas = append(metas, m.Name)
	}
	result = &QueryResult{
		ColumnMetas: metas,
		Result: kr.Result,
	}
	return
}

// 查询Tables列表
func (k *FKKylin) ListTables() (code int, body []byte, err error) {
	if k.ProjectName == ""{
		return 0, nil, fmt.Errorf("Project name shouldn't be empty.")
	}
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/tables_and_columns")
	request.PostArgs().Add("project", k.ProjectName)
	request.Header.SetMethod("GET")
	return k.do(request)
}

// 查询Cube列表
func (k *FKKylin) ListCubes(offset int, limit int) (code int, body []byte, err error) {
	request := k.genBaseRequest()
	request.SetRequestURI(k.Url + "/kylin/api/cubes")
	request.Header.SetMethod("GET")
	request.PostArgs().Add("offset", strconv.Itoa(offset))
	request.PostArgs().Add("limit", strconv.Itoa(limit))
	request.PostArgs().Add("projectName", k.ProjectName)
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