package controllers

// MainController struct
type MainController struct {
	FFBaseController
}

// Get handle HTTP GET request
func (c *MainController) Get() {
	c.FFBaseController.Get()
	c.TplName = "index.html"
}
