package controllers

type PersonController struct {
	FFBaseController
}

func (c *PersonController) Get() {
	c.FFBaseController.Get()

	c.TplName = "person.html"
}
