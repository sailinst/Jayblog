package controllers

import (
	"github.com/astaxie/beego"
	//"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Profile() {
	this.Data["userid"] = "geek"
	this.Data["tag"] = "I am a geek!"
	this.Data["hobby"] = []string{"chess", "code"}

	this.TplName = "user/profile.html"

}
func (this *UserController) Get() {
	this.Ctx.WriteString("appname:" + beego.AppConfig.String("appname") +
		"\nhttpport:" + beego.AppConfig.String("httpport") +
		"\nrunmode:" + beego.AppConfig.String("runmode"))

	/*hp := strconv.Itoa(beego.HttpPort)
	c.Ctx.WriteString("\n\nappname:" + beego.AppName +
		"\nhttpport: " + hp +
		"\nrunmode:" + beego.RunMode)
	*/
}
