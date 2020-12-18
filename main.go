package main

import (
	_ "hello/routers"
	_ "hello/models"
	_ "github.com/go-sql-driver/mysql"
	web "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
)

func init(){
	dbdriver, _ := web.AppConfig.String("dbdriver")
	dbuser, _ := web.AppConfig.String("dbuser")
	dbpassword, _ := web.AppConfig.String("dbpassword")
	dburl, _ := web.AppConfig.String("dburl")
	dbport, _ := web.AppConfig.String("dbport")
	if dbport == "" {
		dbport = "3306"
	}
	dbname, _ := web.AppConfig.String("dbname")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dburl + ":" + dbport + ")/" + dbname + "?charset=utf8" 
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", dbdriver, dsn)
	web.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	web.Run()
}

