package main

import (
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	_ "github.com/inspursoft/itpserver/src/apiserver/routers"
)

const appPath = "conf"

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.LoadAppConfig("ini", filepath.Join(appPath, "app.conf"))
	dao.InitDB()
	beego.Run()
}
