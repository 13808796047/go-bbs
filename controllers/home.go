package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Layout = "layouts/app.html"
	this.TplName = "index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["_header"] = "layouts/_header.html"
	this.LayoutSections["_footer"] = "layouts/_footer.html"

}
