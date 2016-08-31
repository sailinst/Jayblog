package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
		this.Redirect("/", 301)
		return
	}

	this.TplName = "login.html"

}
func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))

	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autologin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/", 301)
	return
}
func checkAccount(ctx *context.Context) bool {
	//ok, err := ctx.Getcookie("username")
	ok, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ok.Value
	ok, _ = ctx.Request.Cookie("password")
	if err != nil {
		return false

	}
	password := ok.Value
	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
