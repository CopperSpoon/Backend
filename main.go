package main

import (
	_ "BBlog/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.BConfig.CopyRequestBody = true
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
