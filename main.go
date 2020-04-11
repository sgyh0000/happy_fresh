package main

import (
	"github.com/astaxie/beego"
	. "happy-fresh/controllers"
	_ "happy-fresh/models"
	_ "happy-fresh/routers"
)

func main() {
	beego.Include(&UserController{}, &CartController{}, &ItemController{}, &MainController{}, &OrderController{}, &CategoryController{})
	beego.Run()
}
