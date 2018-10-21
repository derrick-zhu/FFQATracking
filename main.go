package main

import (
	"FFQATracking/models"
	_ "FFQATracking/routers"

	"github.com/astaxie/beego/logs"

	_ "github.com/gwenn/gosqlite"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	logs.EnableFuncCallDepth(true)
	orm.Debug = false
	models.RegisterDB(false)
}

func main() {
	beego.Run()
}
