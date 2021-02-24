package routers

import (
	"aws-order/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
}
