package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"happy-fresh/models"
	"strconv"
)

func (this *ItemController) URLMapping() {
	this.Mapping("GetItem", this.GetItem)
}

type ItemController struct {
	beego.Controller
}

// @router /item/:id [get]
func (this *ItemController) GetItem() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		log.Error("error", err)
		this.Data["Error"] = "路由失败,请输入正确的商品详情url"
		this.TplName = "error.html"
		return
	}
	var uid = uint64(id)
	var item models.Item
	err = orm.NewOrm().QueryTable(new(models.Item)).Filter("id", uid).One(&item)
	if err != nil {
		log.Error("error", err)
		this.Data["Error"] = "路由失败,请输入正确的商品详情url"
		this.TplName = "error.html"
		return
	}
	this.Data["Item"] = item
	this.TplName = "detail.html"
}
