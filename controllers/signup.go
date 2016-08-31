package controllers

import (
	"Jay_web/Jayblog/models"
	"github.com/astaxie/beego"
)

type Signupcontroller struct {
	beego.Controller
}

func (this *Signupcontroller) Get() {

	adur := this.Input().Get("adur")
	if adur == "add" {
		name := this.Input().Get("name")
		password := this.Input().Get("password")
		email := this.Input().Get("email")
		//不需要判断 name password email 的值，前端已经处理过了。
		err := models.AddUser(name, password, email)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/", 302)
		return
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.Data["IsSignup"] = true

	this.TplName = "signup.html"
}
