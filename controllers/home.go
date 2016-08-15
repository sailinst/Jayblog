package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

}
