package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["name"] = "index.get"
	c.TplName = "index/index.tpl"
}

