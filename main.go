package main

import (
	_ "aws-order/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

