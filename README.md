# FKGoKylin
a simple framework to connect Apache Kylin restful API, write by golang.

# usage

```go
ProjectName := "test_project"
Url := "http://localhost:7070"
UserName := "ADMIN"
Password := "KYLIN"

kylin := FKGoKylin.CreateFKKylin(ProjectName, Url, UserName, Password)\
if kylin == nil{
	return nil, errors.New("create kylin object failed.")
}
_, _, err := kylin.Login()
if err != nil{
	return nil, err
}
code, responseBody, err := kylin.ListTables()
// kylin.ListCubes(0,10)
// kylin.GetCube("")
// kylin.Query("test_table", "column", test_Schema, 0, 10, true)
if err != nil{
	return nil, err
}
fmt.Println(string(responseBody[:]))

result, err := kylin.QueryByStruct("KYLIN_ACCOUNT", []string{"ACCOUNT_ID", "ACCOUNT_COUNTRY"},
		&kylin_table_schema{ACCOUNT_ID:10005647}, 0, 10, true)
if err != nil{
	return nil, err
}
fmt.Println(result.ResultString())

result, err = kylin.QueryBySQL("SELECT COUNT(*) FROM KYLIN_ACCOUNT", 0, 10, true)
if err != nil{
	return nil, err
}
fmt.Println(result.ResultString())
```

** Result **

```
2019/02/21 16:27:56 [INFO][api_simple] kylin_simple.go:104: code = 200 : data = [{"columns":[{"table_NAME":"KYLIN_ACCOUNT","table_SCHEM":"KYLIN_INTERMEDIATE","column_NAME":"ACCOUNT_ID","remarks":null,"type_NAME":"BIGINT","table_CAT":"defaultCatalog","data_TYPE":-5,"column_SIZE":-1,"buffer_LENGTH":-1,"decimal_DIGITS":0,"num_PREC_RADIX":10,"nullable":1,"column_DEF":null,"sql_DATA_TYPE":-1,"sql_DATETIME_SUB":-1,"char_OCTET_LENGTH":-1,"ordinal_POSITION":1,"is_NULLABLE":"YES","scope_CATLOG":null,"scope_SCHEMA":null,"scope_TABLE":null,"source_DATA_TYPE":-1,"is_AUTOINCREMENT":""},{"table_NAME":"KYLIN_SALES","table_SCHEM":"KYLIN_INTERMEDIATE","column_NAME":"OPS_REGION","remarks":null,"type_NAME":"VARCHAR(256) CHARACTER SET \"UTF-16LE\" COLLATE \"UTF-16LE$en_US$primary\"","table_CAT":"defaultCatalog","data_TYPE":12,"column_SIZE":256,"buffer_LENGTH":-1,"decimal_DIGITS":0,"num_PREC_RADIX":10,"nullable":1,"column_DEF":null,"sql_DATA_TYPE":-1,"sql_DATETIME_SUB":-1,"char_OCTET_LENGTH":256,"ordinal_POSITION":10,"is_NULLABLE":"YES","scope_CATLOG":null,"scope_SCHEMA":null,"scope_TABLE":null,"source_DATA_TYPE":-1,"is_AUTOINCREMENT":""}],"table_NAME":"KYLIN_SALES","table_SCHEM":"KYLIN_INTERMEDIATE","self_REFERENCING_COL_NAME":null,"ref_GENERATION":null,"table_TYPE":"TABLE","remarks":null,"type_CAT":null,"type_SCHEM":null,"type_NAME":null,"table_CAT":"defaultCatalog"}]
select ACCOUNT_ID,ACCOUNT_COUNTRY from KYLIN_ACCOUNT where 1=1  and KYLIN_ACCOUNT.ACCOUNT_ID=10005647
2019/02/21 16:27:56 [INFO][api_simple] kylin_simple.go:112: result =  [["10005647","GB"]]
2019/02/21 16:27:56 [INFO][api_simple] kylin_simple.go:119: result =  [["10000"]]
```
