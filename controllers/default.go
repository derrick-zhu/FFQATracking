package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"

	"github.com/astaxie/beego"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get handle HTTP GET request
func (c *MainController) Get() {
	c.Data[constants.Title] = "Farfetch Q&A"
	c.Data[constants.IsLogin] = biz.HadLogin(c.Ctx)
	c.TplName = "index.html"
}
