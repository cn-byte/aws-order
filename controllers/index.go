package controllers

import (
	"aws-order/models/service"
	"aws-order/models/valueobject"
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller

	usersService service.UsersService
}

func (c *IndexController) Get() {
	l := logs.GetLogger()
	list, total, err := c.usersService.QueryListPage(context.Background(), valueobject.UsersListParamVo{Page: 1,
		Size: 10})
	if err != nil {
		logs.Error(1024, "查询出错")
	} else {
		l.Println("查询结果:", list, "数量:", total)
	}
	c.Data["name"] = list
	c.TplName = "index/index.tpl"
}
