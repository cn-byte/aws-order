package main

import (
	"aws-order/database"
	_ "aws-order/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitMysqlConfig()
	beego.Run()
}
