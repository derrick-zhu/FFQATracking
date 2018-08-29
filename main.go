package main

import (
	"FFQATracking/models"
	_ "FFQATracking/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
	models.RegisterDB(true)

}

func main() {

	beego.Run()
}
