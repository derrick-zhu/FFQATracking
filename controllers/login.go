package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/utils"
	"fmt"

	"github.com/astaxie/beego"
)

// LoginController is the page for login
type LoginController struct {
	FFBaseController
}

// Get for handling GET request
func (c *LoginController) Get() {
	c.FFBaseController.Get()

	c.TplName = "login.html"
}

// Signin for handling signin POST request
func (c *LoginController) Signin() {

	email := c.Input().Get(constants.KeyEMAIL)
	pwd := c.Input().Get(constants.KeyPWD)

	beego.Debug(fmt.Sprintf("email: %s, pwd: %s", email, pwd))

	for {
		result, acc := biz.CheckAccount(email, pwd)
		if result == true {

			utils.CookieInstance().Set(c.Ctx, constants.KeyUID, utils.I64toa(int64(acc.ID)), -1)
			utils.CookieInstance().Set(c.Ctx, constants.KeyEMAIL, email, -1)
			utils.CookieInstance().SetSecret(c.Ctx, constants.KeyPWD, pwd, -1)

			utils.MakeRedirectURL(&c.Data, 302, "/", "")
			break
		}

		utils.MakeRedirectURL(&c.Data, 302, "/login/error", "")
		break
	}

	c.ServeJSON()
}

// Signup for handling signup GET request
func (c *LoginController) Signup() {
	c.Redirect("/signup", 302)
	return
}

// Exit for handling exit GET request
func (c *LoginController) Exit() {
	if biz.Logout(c.Ctx) {
		c.Redirect("/", 302)
		return
	}
	c.Redirect("/login", 302)
	return
}
