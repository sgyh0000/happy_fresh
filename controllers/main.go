package controllers

import (
	"github.com/astaxie/beego"
	"happy-fresh/models"
)

func (this *MainController) URLMapping() {
	this.Mapping("Get", this.Get)
	this.Mapping("Search", this.Search)
}

type MainController struct {
	beego.Controller
}

// @router /index [get]
func (this *MainController) Get() {
	this.TplName = "index.html"
}

// @router /search [get]
func (this *MainController) Search() {
	itemName := this.GetString("itemName")
	page, _ := this.GetInt64("page")
	if page == 0 {
		page = 1
	}
	pageSize, _ := this.GetInt64("pageSize")
	if pageSize == 0 {
		pageSize = 20
	}
	pageResult, data := models.FindByKeyWord(itemName, page, pageSize)
	this.Data["page"] = pageResult
	this.Data["item"] = data
	this.TplName = "list.html"
}
