package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

const LAYOUT_PATH = "layout/"

type AppController struct {
	beego.Controller
}

func (this *AppController) setLayout(template string) {
	this.Layout = LAYOUT_PATH + template + ".tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "head.tpl"
	this.LayoutSections["Css"] = "css.tpl"
	this.Data["HomeTitle"], _ = beego.AppConfig.String("homeTitle")
}

