package controllers

type WeatherDemoController struct {
	FFBaseController
}

func (c *WeatherDemoController) Get() {
	c.FFBaseController.Get()

	c.TplName = "weatherDemo.html"
}
