# FKGoKylin
a simple framework to connect Apache Kylin restful API, write by golang.

# usage

```go
ProjectName := "test_project"
Url := "http://localhost:7070"
UserName := "ADMIN"
Password := "KYLIN"

kylin := FKGoKylin.CreateFKKylin(ProjectName, Url, UserName, Password)
code, responseBody, err := kylin.ListTables()
// kylin.ListCubes(0,10)
// kylin.GetCube("")
// kylin.Query("test_table", "column", test_Schema, 0, 10, true)
if err != nil{
  return
}
fmt.Println(string(responseBody[:]))
```
