package routers

import (
	"github.com/astaxie/beego"
	"go-bbs/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
