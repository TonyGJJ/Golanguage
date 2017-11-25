package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// (this *MainController) ...
func (this *MainController)Get ()  {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
