package render

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

type col struct {
	ColumnName string
	//DataType      string
	//ColumnType    string
	ColumnComment string
	ModelProp     string
	ModelType     string
	ModelTag      string
	//ColumnKey     string
	//ColumnType   string
	SwaggerType   string
	SwaggerFormat string
}

func newCol(cName, dType, cComment, cKey, cType string) col {
	modelProperty := strcase.ToCamel(cName)
	modelType := "NoneType"
	swgType, swgFormat := "", ""
	switch dType {
	case "varchar", "longtext", "char", "enum", "set", "mediumtext", "json", "text":
		modelType = "string"
		swgFormat, swgType = "string", "string"
	case "bigint", "int", "tinyint", "smallint":
		modelType = "int"
		swgFormat, swgType = "int64", "integer"

	case "decimal", "float", "double":
		swgFormat, swgType = "float", "number"
		modelType = "float64"
	case "blob":
		swgFormat, swgType = "binary", "string"
		modelType = "*[]byte"
	case "time", "datetime", "timestamp":
		swgFormat, swgType = "date-time", "string"
		modelType = "*time.Time"
	}
	if strings.Contains(cType, "unsigned") {
		modelType = "uint"
		swgFormat, swgType = "int64", "integer"

	}
	if cKey == "PRI" {
		modelType = "uint"
		swgFormat, swgType = "int64", "integer"

	}
	//Content   string     `form:"content" json:"content,omitempty" comment:""`
	pt := fmt.Sprintf("%s   %s", modelProperty, modelType)
	sql := ""
	if cType != "" {
		sql = fmt.Sprintf("%s%s", sql, cType)
	}
	if cKey != "" {
		sql = fmt.Sprintf("%s,%s", sql, cKey)
	}
	if dType == "json" {
		cComment = fmt.Sprintf("must a string can unmarsh to an Object. %s", cComment)
	}
	tag := fmt.Sprintf(`form:"%s" json:"%s,omitempty" comment:"%s" sql:"%s"`, cName, cName, cComment, sql)
	modelTag := fmt.Sprintf("%s     `%s`", pt, tag)
	return col{cName, cComment, modelProperty, modelType, modelTag, swgType, swgFormat}
}
