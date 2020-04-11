package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"happy-fresh/models"
	"strings"
)

var (
	copyright = beego.AppConfig.String("footer_copyright")
	detail    = beego.AppConfig.String("footer_detail")
)

var LoginFilter = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Request.RequestURI, "/static") {
		return
	}
	_, ok := ctx.Input.Session("User").(models.User)
	if !ok && ctx.Request.RequestURI != "/" && ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/register" {
		ctx.Redirect(302, "/")
		_ = ctx.Input.CruSession.Set("referer", ctx.Request.RequestURI)
	} else if ok && ctx.Request.RequestURI == "/" {
		ctx.Redirect(302, "/index")
	} else if ok && ctx.Request.RequestURI == "/login" {
		ctx.Redirect(302, "/index")
	}
}

var PublicDataFilter = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Request.RequestURI, "/static") {
		return
	}
	user, ok := ctx.Input.Session("User").(models.User)
	if !ok {
		return
	}
	categories := models.SelectAll()

	ctx.Input.SetData("Copyright", copyright)
	ctx.Input.SetData("Detail", detail)

	ctx.Input.SetData("Name", user.RealName)
	ctx.Input.SetData("Categories", categories)
	ctx.Input.SetData("CartCount", models.FindCountByUser(user.UserName))
	ctx.Input.SetData("NewItems", models.FindNewItems())
}

var ClueFilter = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Request.RequestURI, "/static") {
		return
	}
	user, ok := ctx.Input.Session("User").(models.User)
	if !ok {
		return
	}
	uri := ctx.Request.RequestURI
	var clue = &models.Clue{
		Uri:      uri,
		UserName: user.UserName,
	}
	_, e := orm.NewOrm().Insert(clue)
	fmt.Print(e)
}

func init() {
	beego.InsertFilter("/*", beego.BeforeExec, PublicDataFilter)
	beego.InsertFilter("/*", beego.BeforeRouter, LoginFilter)
	beego.InsertFilter("/*", beego.BeforeRouter, ClueFilter)
}
