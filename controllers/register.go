package controllers

import "FFQATracking/constants"

// RegisterController class
type RegisterController struct {
	FFBaseController
}

// Get for handling GET request
func (c *RegisterController) Get() {
	c.FFBaseController.Get()

	permissionRules := []string{
		"hello",
		"world",
	}
	c.Data[constants.KeyPermissions] = permissionRules

	c.TplName = "register.html"
}
