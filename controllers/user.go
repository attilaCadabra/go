package controllers

type UserController struct {
	AppController
}

func (this *UserController) Home() {
	user := this.GetSession("AuthUser")
	if user != nil {
		this.Data["AuthUser"] = user
	}
	this.setLayout("default")
	this.TplName = "home.tpl"
}
