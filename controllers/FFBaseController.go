package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/utils"

	"github.com/astaxie/beego"
)

// FFBaseController is the basic *Controller* class
type FFBaseController struct {
	beego.Controller
}

// Get is the GET method for handler default Get request
func (c *FFBaseController) Get() {
	c.Data[constants.Title] = "Farfetch Q&A Tracking"
	c.Data[constants.IsLogin] = biz.HadLogin(c.Ctx)

	email := utils.CookieInstance().Get(c.Ctx, constants.KeyEMAIL)
	uname := utils.CookieInstance().Get(c.Ctx, constants.KeyUNAME)

	if len(email) > 0 {
		c.Data[constants.LoggedInAccount] = email
	} else if len(uname) > 0 {
		c.Data[constants.LoggedInAccount] = uname
	}
}

// Post is the POST method for handler default Post request
func (c *FFBaseController) Post() {

}
