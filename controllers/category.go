package controllers

import (
	"github.com/astaxie/beego"
	"happy-fresh/models"
	"strconv"
)

func (this *CategoryController) URLMapping() {
	this.Mapping("Category", this.Category)
}

type CategoryController struct {
	beego.Controller
}

// @router /category/:cid [get]
func (this *CategoryController) Category() {
	cid, err := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	if err != nil {
		log.Error("error", err)
		this.Data["Error"] = "路由失败,请输入正确的商品详情url"
		this.TplName = "error.html"
		return
	}
	page, _ := this.GetInt64("page")
	if page == 0 {
		page = 1
	}
	pageSize, _ := this.GetInt64("pageSize")
	if pageSize == 0 {
		pageSize = 20
	}
	pageResult, data := models.FindPageByCategoryId(models.Pk(cid), page, pageSize)
	this.Data["page"] = pageResult
	this.Data["item"] = data
	this.TplName = "list.html"
}
