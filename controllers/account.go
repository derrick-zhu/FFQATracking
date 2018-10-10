package controllers

// AccountController account page controller
type AccountController struct {
	FFBaseController
}

// Get handle AccountController's GET http request
func (c *AccountController) Get() {
	c.FFBaseController.Get()

	c.TplName = "account.html"
}
