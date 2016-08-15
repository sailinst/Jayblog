package controllers

import (
	"github.com/astaxie/beego"
)

type Signupcontroller struct {
	beego.Controller
}

func (this *Signupcontroller) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.Data["IsSignup"] = true

	this.TplName = "signup.html"
}
