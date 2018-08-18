package main

import (
	"github.com/astaxie/beego"
	"odata-batch/controller"
)

func main() {
	beego.AddNamespace(
		beego.NewNamespace(
			"calculator",
			beego.NSRouter("/$batch", &controller.Calculator{}, "post:Batch"),
			beego.NSRouter("/Actions/Calculator.Reset", &controller.Calculator{}, "post:Reset"),
			beego.NSRouter("/Actions/Calculator.Add", &controller.Calculator{}, "post:Add"),
			beego.NSRouter("/Actions/Calculator.Sub", &controller.Calculator{}, "post:Sub"),
			beego.NSRouter("/Actions/Calculator.Mul", &controller.Calculator{}, "post:Mul"),
			beego.NSRouter("/Actions/Calculator.Div", &controller.Calculator{}, "post:Div"),
		),
	)
	beego.BConfig.Listen.HTTPPort = 3000
	// beego.BConfig.CopyRequestBody = true
	beego.Run()
}