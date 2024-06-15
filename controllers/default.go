package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "localhost"
	c.Data["Email"] = "phpclub@gmail.com"
	c.TplName = "index.tpl"
}
