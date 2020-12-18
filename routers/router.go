package routers

import (
	"hello/controllers"
	web "github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.AuthController{}, "get:Login")
	web.Router("/", &controllers.AuthController{}, "post:DoLogin")
	web.Router("/user", &controllers.UserController{}, "get:Home")
}
