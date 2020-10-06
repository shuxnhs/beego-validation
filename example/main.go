package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.AutoRouter(&ExampleController{})
	beego.BConfig.Listen.HTTPPort = 8888
	beego.Run()
}
