package controllers

import (
	"github.com/astaxie/beego"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get handle HTTP GET request
func (c *MainController) Get() {
	c.Data["Title"] = "Farfetch Q&A"
	c.TplName = "index.html"
}
