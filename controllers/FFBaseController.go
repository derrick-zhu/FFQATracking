package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"

	"github.com/astaxie/beego"
)

type FFBaseController struct {
	beego.Controller
}

func (c *FFBaseController) Get() {
	c.Data[constants.Title] = "Farfetch Q&A"
	c.Data[constants.IsLogin] = biz.HadLogin(c.Ctx)
}
