package controllers

import (
	"hello/models"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
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
	password := this.GetString("hash")
	user := AuthUser{UserName: userName}

	o := orm.NewOrm()
	err := o.Read(&user, "UserName")

	if err != nil {
		this.Data["Error"] = "ユーザー名またはパスワードが間違っています"
		defer this.Login()
	}else{
		er := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
		if er == nil{
			this.SetSession("AuthUser", user)
			this.Ctx.Redirect(302, "/user")
		}else{
			this.Data["Error"] = "ユーザー名またはパスワードが間違っています"
			defer this.Login()
		}
	}
}