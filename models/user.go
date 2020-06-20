package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
