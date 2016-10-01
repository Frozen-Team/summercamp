package main

import (
	_ "bitbucket.org/SummerCampDev/summercamp/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"
	}
	beego.Run()
}
