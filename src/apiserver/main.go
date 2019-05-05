package main

import (
	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	_ "github.com/inspursoft/itpserver/src/apiserver/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	dao.InitDB()
	beego.Run()
}
