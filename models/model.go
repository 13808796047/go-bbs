package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%v:%v@%v(%v)/%v?charset=utf8",
			beego.AppConfig.String("DB_USERNAME"),
			beego.AppConfig.String("DB_PASSWORD"),
			beego.AppConfig.String("DB_PORT"),
			beego.AppConfig.String("DB_HOST"),
			beego.AppConfig.String("DB_DATABASE")))
}
