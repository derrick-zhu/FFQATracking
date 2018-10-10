package controllers

// PersonController user's personal profile page
type PersonController struct {
	FFBaseController
}

// Get handle PersonController's GET http request
func (c *PersonController) Get() {
	c.FFBaseController.Get()

	c.TplName = "person.html"
}
