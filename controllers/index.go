package controllers

import (
	"aws-order/models/entity"
	"aws-order/models/service"
	"aws-order/models/valueobject"
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gogf/gf/util/gconv"
)

type IndexController struct {
	beego.Controller

	usersService service.UsersService
}

func (c *IndexController) Get() {
	id, _ := c.GetInt32(":id")
	l := logs.GetLogger()
	detail, err := c.usersService.QueryDetail(context.Background(), entity.Users{Id: id})
	if err != nil {
		logs.Error(1024, err)
	} else {
		l.Println("查询结果:", detail)
	}
	c.Data["name"] = detail
	c.TplName = "index/index.tpl"
}

// GetAll
func (c *IndexController) GetAll() {
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

/**
{
 "username":"占山",
 "password":"55555",
 "id":1
}
*/
//  Save json数据
func (c *IndexController) Save() {
	var usersSaveParamVo valueobject.UsersSaveParamVo
	_ = gconv.Struct(c.Ctx.Input.RequestBody, &usersSaveParamVo)
	err := c.usersService.Save(context.Background(), usersSaveParamVo)
	if err != nil {
		logs.Error(1024, "查询出错")
	}
	if usersSaveParamVo.Id > 0 {
		c.Data["name"] = "修改成功"
	} else {
		c.Data["name"] = "创建成功"
	}
	c.TplName = "index/index.tpl"
}
