package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"

	"github.com/astaxie/beego"
)

// FFBaseController is the basic *Controller* class
type FFBaseController struct {
	beego.Controller
}

// Get is the GET method for handler default Get request
func (c *FFBaseController) Get() {
	c.Data[constants.Title] = "Farfetch Q&A"
	c.Data[constants.IsLogin] = biz.HadLogin(c.Ctx)
}

// Post is the POST method for handler default Post request
func (c *FFBaseController) Post() {

}
