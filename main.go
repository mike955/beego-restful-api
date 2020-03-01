package main

import (
	_ "beego-restful-api/routers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:3564423@/beego_api?charset=utf8mb4")
}


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	beego.Run()
}
