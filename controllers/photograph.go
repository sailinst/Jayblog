package controllers

import (
	"github.com/astaxie/beego"
)

type PhotographController struct {
	beego.Controller
}

func (this *PhotographController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.Data["IsPhotograph"] = true

	this.TplName = "photograph.html"
}
