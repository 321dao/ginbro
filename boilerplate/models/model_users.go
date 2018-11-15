package models

import (
	"errors"
	"time"
)

type User struct {
	Id uint `form:"id" json:"id,omitempty" comment:""`

	Username string `form:"username" json:"username,omitempty" comment:"昵称/登陆用户名"`

	FullName string `form:"full_name" json:"full_name,omitempty" comment:"真实姓名"`

	Email string `form:"email" json:"email,omitempty" comment:"邮箱"`

	Mobile string `form:"mobile" json:"mobile,omitempty" comment:"手机号码"`

	Password string `form:"password" json:"password,omitempty" comment:"密码"`

	RoleId uint `form:"role_id" json:"role_id,omitempty" comment:"角色ID:2-超级用户,4-普通用户"`

	Status uint `form:"status" json:"status,omitempty" comment:"状态: 1-正常,2-禁用/删除"`

	Avatar string `form:"avatar" json:"avatar,omitempty" comment:"用户头像"`

	Remark string `form:"remark" json:"remark,omitempty" comment:"备注"`

	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" comment:""`

	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" comment:""`
}

func (m *User) TableName() string {
	return "users"
}
func (m *User) One() (one *User, err error) {
	one = &User{}
	err = crudOne(m, one)
	return
}

func (m *User) All(q *PaginationQuery) (list *[]User, total uint, err error) {
	list = &[]User{}
	total, err = crudAll(m, q, list)
	return
}

func (m *User) Update() (err error) {
	where := User{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *User) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *User) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
