package main

import (
	"FFQATracking/models"
	_ "FFQATracking/routers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
	models.RegisterDB(false)

}

func main() {

	beego.Run()
}
