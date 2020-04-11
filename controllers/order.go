package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"happy-fresh/models"
	"strconv"
	"strings"
)

var log = logs.GetBeeLogger()

func (this *OrderController) URLMapping() {
	this.Mapping("AddOrder", this.AddOrder)
}

type OrderController struct {
	beego.Controller
}

// @router /order/add [get]
func (this *OrderController) AddOrder() {
	cartIdStr := this.GetString("cartIds")
	var cartIds []models.Pk

	for _, str := range strings.Split(cartIdStr, ",") {
		id, _ := strconv.Atoi(str)
		cartId := models.Pk(id)
		cartIds = append(cartIds, cartId)
	}

	user := this.GetSession("User").(models.User)

	models.AddOrder(cartIds, user.UserName)
	this.Redirect("/user_center/order", 302)
}
