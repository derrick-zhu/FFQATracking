package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/models"

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

	var acc *models.AccountModel
	var err error
	if acc, err = biz.CurrentAccount(c.Ctx); err != nil {
		beego.Error(err)
	}

	c.Data[constants.AccountData] = acc
}

// Post is the POST method for handler default Post request
func (c *FFBaseController) Post() {

}
