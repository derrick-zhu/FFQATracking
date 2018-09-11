package controllers

import "FFQATracking/constants"

// MainController struct
type MainController struct {
	FFBaseController
}

// Get handle HTTP GET request
func (c *MainController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "Farfetch Q&A Tracking"
	c.Data[constants.KeyIsHome] = 1

	c.TplName = "index.html"
}
