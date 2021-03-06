package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/models"
	"FFQATracking/utils"

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

	email := c.Input().Get(constants.KeyEMAIL)
	pwd := c.Input().Get(constants.KeyPWD)

	beego.Info("Get ready for register account email: " + email + ", pwd: " + pwd)

	for {
		if biz.HasAccountIfNot(email) == true {

			result, err = biz.Login(c.Ctx, email, pwd)
			if err != nil || result == false {

				beego.Error(err)
				utils.MakeRedirectURL(&c.Data, 302, "#", "")
				break
			}

			utils.MakeRedirectURL(&c.Data, 302, "/", "")
			break
		}

		beego.Info("Do registering ...")
		result, _, err = biz.Register(email, pwd, models.RuleUser)
		if err != nil || result == false {

			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		beego.Info("Do login ...")
		result, err = biz.Login(c.Ctx, email, pwd)
		if err != nil || result == false {

			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		beego.Info("Finish register progress ...")
		utils.MakeRedirectURL(&c.Data, 302, "/", "")

		break
	}

	c.ServeJSON()
	return
}
