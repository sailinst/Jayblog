package controllers

import (
	"Jay_web/Jayblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	//检查是否有操作
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
