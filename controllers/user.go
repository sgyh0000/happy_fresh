package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"happy-fresh/models"
	"regexp"
	"strconv"
	"strings"
)

func (this *UserController) URLMapping() {
	this.Mapping("Login", this.Login)
	this.Mapping("Get", this.Get)
	this.Mapping("RegisterRouter", this.RegisterRouter)
	this.Mapping("Register", this.Register)
	this.Mapping("UserInfo", this.UserInfo)
	this.Mapping("UserSite", this.UserSite)
	this.Mapping("EditDelivery", this.EditDelivery)
}

type UserController struct {
	beego.Controller
}

// @router /login [post]
func (this *UserController) Login() {
	postUser := models.User{}
	err := this.ParseForm(&postUser)
	if err != nil {
		// error handle
		this.TplName = "error.html"
		return
	}

	rows := models.FindCountByUserNameAndPassword(postUser.UserName, postUser.Password)
	if rows <= 0 {
		// error handle
		this.TplName = "login.html"
		this.Data["Error"] = true
	} else {
		// set user session
		user, _ := models.FindUserByUserName(postUser.UserName)
		this.SetSession("User", user)

		// auto Redirect to the uri that visit before
		referer := this.GetSession("referer").(string)
		if referer != "" {
			this.Redirect(referer, 302)
			_ = this.CruSession.Delete("referer")
			return
		}

		this.Redirect("/index", 302)
	}
}

// @router / [get]
func (this *UserController) Get() {
	this.TplName = "login.html"
}

// @router /register [get]
func (this *UserController) RegisterRouter() {
	this.TplName = "register.html"
}

// @router /register/submit [post]
func (this *UserController) Register() {
	postUser := models.User{}
	_ = this.ParseForm(&postUser)
	id := models.InsetUser(&postUser)
	if id > 0 {
		postUser.Id = models.Pk(id)
		this.Redirect("/", 302)
		this.SetSession("User", postUser)
	} else {
		this.Data["Error"] = errors.New("注册失败")
		this.TplName = "error.html"
	}
}

// @router /user_center/info [get]
func (this *UserController) UserInfo() {
	this.TplName = "user_center_info.html"
	user := this.GetSession("User").(models.User)
	this.Data["Delivery"] = models.FindByUserName(user.UserName)
	this.Data["CurrentPathName"] = "用户中心"

	clues := models.FindCluesByUserName(user.UserName)
	// no clues of item
	if len(clues) == 0 {
		return
	}

	var itemIds []models.Pk
	for _, clue := range clues {
		itemId, _ := strconv.Atoi(strings.Split(clue.Uri, "/item/")[1])
		itemIds = append(itemIds, models.Pk(itemId))
	}
	this.Data["DiyItems"], _ = models.FindItemByIds(itemIds)

}

// @router /user_center/site [get]
func (this *UserController) UserSite() {
	this.TplName = "user_center_site.html"
	user := this.GetSession("User").(models.User)

	delivery := models.FindByUserName(user.UserName)
	telephoneBytes := []byte(delivery.Telephone)
	for i, _ := range telephoneBytes {
		if i >= 3 && i <= 6 {
			telephoneBytes[i] = '*'
		}
	}
	delivery.Telephone = string(telephoneBytes)
	this.Data["Delivery"] = delivery
	this.Data["CurrentPathName"] = "用户中心"
}

// @router /user_center/order [get]
func (this *UserController) UserOrder() {
	user := this.GetSession("User").(models.User)

	orders := models.FindOrderByUserName(user.UserName)
	this.Data["Orders"] = orders

	this.TplName = "user_center_order.html"
}

// @router /delivery/edit [post]
func (this *UserController) EditDelivery() {
	delivery := models.Delivery{}
	_ = this.ParseForm(&delivery)

	errorMap := checkDelivery(delivery)

	if len(errorMap) > 0 {
		this.Data["Error"] = errorMap
	} else {
		userName := this.GetSession("User").(models.User).UserName
		delivery.UserName = userName
		models.InsertOrUpdateDelivery(delivery)
	}
	this.Redirect("/user_center/site", 302)
}

func checkDelivery(delivery models.Delivery) map[string]string {
	errorMap := map[string]string{}

	if len(delivery.RealName) == 0 {
		errorMap["RealName"] = realNameError
	}

	rgx := regexp.MustCompile(telReg)

	if !rgx.MatchString(delivery.Telephone) {
		errorMap["Telephone"] = telephoneError
	}

	if len(delivery.Code) != 6 {
		errorMap["Code"] = codeError
	}

	if len(delivery.Address) == 0 {
		errorMap["Address"] = addressError
	}

	return errorMap
}

const (
	realNameError  = "请填写姓名"
	telephoneError = "请填写正确的手机号"
	codeError      = "请填写正确的邮编"
	addressError   = "请填写收货地址"

	telReg = `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
)
