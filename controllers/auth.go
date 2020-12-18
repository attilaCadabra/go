package controllers

import (
	"hello/models"
	"github.com/beego/beego/v2/client/orm"
)

type AuthUser = models.User

type AuthController struct {
	AppController
}

func (this *AuthController) Login() {
	this.setLayout("login")
	this.TplName = "login.tpl"
}

func (this *AuthController) DoLogin() {
	userName := this.GetString("user_name")
	hash := this.GetString("hash")
	user := AuthUser{UserName: userName, Hash: hash}

	o := orm.NewOrm()
	err := o.Read(&user, "UserName", "Hash")

	if err != nil {
		this.Data["Error"] = "ユーザー名またはパスワードが間違っています"
		this.Login()
	}else{
		this.Ctx.Redirect(302, "/user")
	}
}