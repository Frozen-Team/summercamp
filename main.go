package main

import (
	_ "github.com/Frozen-Team/summercamp/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	fmt.Println("Application version:", version)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"
	}
	beego.Run()
}
