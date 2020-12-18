package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id int
	UserName string
	Hash string
}

func init(){
	orm.RegisterModel(new(User))
}