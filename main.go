package main

import (
	"Jay_web/Jayblog/controllers"
	"Jay_web/Jayblog/models"
	_ "Jay_web/Jayblog/routers"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
var readtimes int

func checkErr(funcname string, err error) {
	if err != nil {
		fmt.Printf("%s() is failed:%s\n", funcname, err)
		os.Exit(-1)
	}
}
*/
func init() {
	models.RegisterDB()
}
func main() {

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/photograph", &controllers.PhotographController{})
	beego.Run()

}
