package routers

import (
	"hello/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.AuthController{}, "get:Login")
	beego.Router("/", &controllers.AuthController{}, "post:DoLogin")
	beego.Router("/user", &controllers.UserController{}, "get:Home")
}
