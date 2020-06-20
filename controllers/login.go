package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-bbs/models"
	"html/template"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Layout = "layouts/app.html"
	this.TplName = "auth/login.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["_header"] = "layouts/_header.html"
	this.LayoutSections["_footer"] = "layouts/_footer.html"

}
func (this *LoginController) Post() {
	u := models.User{}

	if err := this.ParseForm(&u); err != nil {
		this.Redirect("/login", 302)
	}
	fmt.Println(u)

}
