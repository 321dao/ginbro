package render

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type schema struct {
	ResourceName   string
	HandlerName    string
	TableName      string
	ModelName      string
	Columns        []col
	ProjectPath    string
	ProjectPackage string
}

func newTable(tableName, projectPath, projectPackage string) *schema {
	mName := strcase.ToCamel(tableName)
	mName = inflection.Singular(mName)
	rName := strcase.ToKebab(mName)
	hName := strcase.ToLowerCamel(mName)
	return &schema{ModelName: mName, TableName: tableName, ResourceName: rName, HandlerName: hName, ProjectPath: projectPath, ProjectPackage: projectPackage}
}

func (t *schema) generateModel() error {
	fileName := fmt.Sprintf("models/model_%s.go", t.TableName)
	return parseTmpl("models.template.tpl", t.ProjectPackage, fileName, t)
}

func (t *schema) generateHandler() error {
	fileName := fmt.Sprintf("handlers/handler_%s.go", t.TableName)
	return parseTmpl("handlers.template.tpl", t.ProjectPackage, fileName, t)
}
