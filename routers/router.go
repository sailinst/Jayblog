package routers

import (
	"Jay_web/Jayblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/photograph", &controllers.PhotographController{})
	beego.Router("/signup", &controllers.Signupcontroller{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("/api/user/profile", &controllers.UserController{}, `get:API_Profile`)

}
