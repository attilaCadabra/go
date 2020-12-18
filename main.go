package main

import (
	_ "hello/routers"
	_ "hello/models"
	_ "github.com/go-sql-driver/mysql"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
)

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.175.27:3306)/go?charset=utf8")
}

func main() {
	beego.Run()
}

