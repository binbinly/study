package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"lottery/controllers"
)

func main() {
	controllers.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()
	beego.Run()
}
