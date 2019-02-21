package FKGoKylin

import (
	"encoding/json"
	"fmt"
)

type QueryCondition struct {
	SQL           string `json:"sql"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	ProjectName   string `json:"projectName"`
}

func (qc *QueryCondition) ToBytes() (result []byte) {
	var err error
	result, err = json.Marshal(qc)
	if err != nil {
		fmt.Println("Convert query condition to json error.", err)
		return nil
	}
	return
}

type Column struct {
	IsNullable         int    `json:"isNullable"`
	DisplaySize        int    `json:"displaySize"`
	Label              string `json:"label"`
	Name               string `json:"name"`
	SchemaName         string `json:"schemaName"`
	CatelogName        string `json:"catelogName"`
	TableName          string `json:"tableName"`
	Precision          int    `json:"precision"`
	Scale              int    `json:"scale"`
	ColumnType         int    `json:"columnType"`
	ColumnTypeName     string `json:"columnTypeName"`
	ReadOnly           bool   `json:"readOnly"`
	AutoIncrement      bool   `json:"autoIncrement"`
	CaseSensitive      bool   `json:"caseSensitive"`
	Searchable         bool   `json:"searchable"`
	Currency           bool   `json:"currency"`
	Signed             bool   `json:"signed"`
	DefinitelyWritable bool   `json:"definitelyWritable"`
	Writable           bool   `json:"writable"`
}

type KylinResult struct {
	ColumnMetas       []*Column     `json:"columnMetas"`
	Result            []interface{} `json:"results"`
	Cube              string        `json:"cube"`
	AffectedRowCount  int           `json:"affectedRowCount"`
	IsException       bool          `json:"isException"`
	ExceptionMessage  string        `json:"exceptionMessage"`
	Duration          int           `json:"duration"`
	TotalScanCount    int           `json:"totalScanCount"`
	HitExceptionCache bool          `json:"hitExceptionCache"`
	StorageCacheUsed  bool          `json:"storageCacheUsed"`
	Partial           bool          `json:"partial"`
}

type QueryResult struct{
	ColumnMetas []string      `json:"columnMetas"`
	Result      []interface{} `json:"results"`
}
