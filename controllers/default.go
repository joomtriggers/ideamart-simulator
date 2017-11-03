package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "TEsting Out"
	c.Data["Email"] = "gnanakeethan@gmail.com"
	c.TplName = "index.tpl"
}
