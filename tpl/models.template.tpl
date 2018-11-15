package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

type {{.ModelName}} struct {
	{{range .Columns}}
	{{.ModelTag}}{{end}}
}

func (m *{{.ModelName}}) TableName() string {
	return "{{.TableName}}"
}

func (m *{{.ModelName}}) One() (one *{{.ModelName}}, err error) {
	one = &{{.ModelName}}{}
	err = crudOne(m, one)
	return
}

func (m *{{.ModelName}}) All(q *PaginationQuery) (list *[]{{.ModelName}}, total uint, err error) {
	list = &[]{{.ModelName}}{}
	total, err = crudAll(m, q, list)
	return
}

func (m *{{.ModelName}}) Update() (err error) {
	where := {{.ModelName}}{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *{{.ModelName}}) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *{{.ModelName}}) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
