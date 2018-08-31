package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

// RegisterController class
type RegisterController struct {
	FFBaseController
}

// Get for handling GET request
func (c *RegisterController) Get() {
	c.FFBaseController.Get()

	c.TplName = "register.html"
}

// Post for handling POST request
func (c *RegisterController) Post() {
	c.FFBaseController.Post()

	var result bool
	var err error

	uname := c.Input().Get(constants.KeyUNAME)
	pwd := c.Input().Get(constants.KeyPWD)

	beego.Info("Get ready for register account uname: " + uname + ", pwd: " + pwd)

	if biz.HasAccountIfNot(uname) == true {
		result, err = biz.Login(c.Ctx, uname, pwd)
		if err != nil || result == false {
			beego.Error(err)
			c.Redirect("#", 302)
			return
		}
		c.Redirect("/", 302)
		return
	}

	beego.Info("Do registering ...")
	result, _, err = biz.Register(uname, pwd, models.RuleUser)
	if err != nil || result == false {
		beego.Error(err)
		c.Redirect("#", 302)
		return
	}

	beego.Info("Do login ...")
	result, err = biz.Login(c.Ctx, uname, pwd)
	if err != nil || result == false {
		beego.Error(err)
		c.Redirect("#", 302)
		return
	}

	beego.Info("Finish register progress ...")
	c.Redirect("/", 302)
	return
}
