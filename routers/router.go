package routers

import (
	"aws-order/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index/:id([0-9]+)", &controllers.IndexController{})
	beego.Router("/index/all", &controllers.IndexController{}, "get:GetAll")
	beego.Router("/index/save", &controllers.IndexController{}, "post:Save")
}
