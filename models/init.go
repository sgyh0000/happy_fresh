package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Pk uint64
type Yuan float32

var log = logs.GetBeeLogger()
var o orm.Ormer

type mysqlConfig struct {
	user string
	pass string
	url  string
	db   string
	port string
}

var config = mysqlConfig{
	user: beego.AppConfig.String("mysql_user"),
	pass: beego.AppConfig.String("mysql_pass"),
	url:  beego.AppConfig.String("mysql_urls"),
	db:   beego.AppConfig.String("mysql_db"),
	port: beego.AppConfig.String("mysql_port"),
}

func init() {
	registerModels(new(Item), new(Category), new(Order), new(User), new(Cart), new(Delivery), new(Clue), new(OrderDetail))
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		log.Error("dao - init: ", err)
	}
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&loc=Local", config.user, config.pass, config.url, config.port, config.db)
	err = orm.RegisterDataBase("default", "mysql", dataSource, 30)
	if err != nil {
		log.Error("dao - init", err)
	}

	_ = orm.RunSyncdb("default", false, true)

	o = orm.NewOrm()
}

func registerModels(models ...interface{}) {
	orm.RegisterModel(models...)
}
