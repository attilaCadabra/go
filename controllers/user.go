package controllers

type UserController struct {
	AppController
}

func (this *UserController) Home() {
	this.setLayout("default")
	this.TplName = "home.tpl"
}
