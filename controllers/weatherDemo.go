package controllers

import (
	"FFQATracking/utils"

	"github.com/astaxie/beego"
)

type WeatherDemoController struct {
	FFBaseController
}

func (c *WeatherDemoController) Get() {
	c.FFBaseController.Get()

	beego.Info(c)

	c.TplName = "weatherDemo.html"
}

func (c *WeatherDemoController) Post() {

	beego.Info(c)

	utils.MakeRedirectURL(&c.Data, 302, "#", "in WeatherDemoController")
	c.ServeJSON()
}
