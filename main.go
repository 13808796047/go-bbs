package main

import (
	"github.com/astaxie/beego"
	_ "go-bbs/models"
	_ "go-bbs/routers"
)

func main() {
	beego.Run()
}
