package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"happy-fresh/models"
	"strconv"
)

func (this *CartController) URLMapping() {
	this.Mapping("Add", this.Add)
	this.Mapping("Get", this.Get)
	this.Mapping("EditCart", this.EditCart)
	this.Mapping("DeleteById", this.DeleteById)
}

type CartController struct {
	beego.Controller
}

// @router /cart [get]
func (this *CartController) Get() {
	user := this.GetSession("User").(models.User)
	cart := models.FindByUser(user.UserName)
	this.Data["Carts"] = cart
	this.TplName = "cart.html"
}

// @router /cart/add [get]
func (this *CartController) Add() {
	id, err := strconv.Atoi(this.GetString("itemId"))
	if err != nil {
		this.TplName = "error.html"
		return
	}
	num, err := strconv.Atoi(this.GetString("num"))
	uid := models.Pk(id)
	user := this.GetSession("User").(models.User)
	cart := &models.Cart{
		UserName: user.UserName,
		ItemId:   uid,
		Num:      uint64(num),
	}
	_, _ = orm.NewOrm().Insert(cart)
	cartCount := models.FindCountByUser(user.UserName)
	result := Result{
		Code:    200,
		Success: true,
		Msg:     "success",
		Err:     nil,
		Data:    cartCount,
	}
	data, err := json.Marshal(&result)
	if err != nil {
		this.TplName = "error.html"
		this.Data["Error"] = err
		return
	}
	this.Data["json"] = string(data)
	this.ServeJSON()
}

// @router /cart/edit/:id [get]
func (this *CartController) EditCart() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	cartId := models.Pk(id)

	op := this.GetString("op")
	result := &Result{
		Code:    200,
		Success: true,
		Msg:     "操作成功",
		Err:     nil,
		Data:    nil,
	}
	if op == "add" {
		result.Data = models.AddItem(cartId)
	} else if op == "minus" {
		result.Data = models.MinusItem(cartId)
	} else {
		result.Success = false
		result.Code = 403
		result.Msg = "不支持的操作"
	}
	data, _ := json.Marshal(result)
	this.Data["json"] = string(data)
	this.ServeJSON()
}

// @router /cart/delete/:id [get]
func (this *CartController) DeleteById() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	cartId := models.Pk(id)
	models.DeleteByCartId(cartId)
	result, _ := json.Marshal(&Result{
		Code:    200,
		Success: true,
		Msg:     "删除成功",
		Err:     nil,
		Data:    nil,
	})
	this.Data["json"] = string(result)
	this.ServeJSON()
}
