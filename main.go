package main

import (
	_ "hello/routers"
	_ "hello/models"
	_ "github.com/go-sql-driver/mysql"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
)

func init(){
	dbdriver, _ := beego.AppConfig.String("dbdriver")
	dbuser, _ := beego.AppConfig.String("dbuser")
	dbpassword, _ := beego.AppConfig.String("dbpassword")
	dburl, _ := beego.AppConfig.String("dburl")
	dbport, _ := beego.AppConfig.String("dbport")
	if dbport == "" {
		dbport = "3306"
	}
	dbname, _ := beego.AppConfig.String("dbname")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dburl + ":" + dbport + ")/" + dbname + "?charset=utf8" 
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", dbdriver, dsn)
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

