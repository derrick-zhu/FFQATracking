package controllers

type AccountController struct {
	FFBaseController
}

func (c *AccountController) Get() {
	c.FFBaseController.Get()

	c.TplName = "account.html"
}
