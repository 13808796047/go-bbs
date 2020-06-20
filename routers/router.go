package routers

import (
	"github.com/astaxie/beego"
	"go-bbs/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
}
