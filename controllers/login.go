package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/utils"
)

// LoginController is the page for login
type LoginController struct {
	FFBaseController
}

// Get for handling GET request
func (c *LoginController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "Farfetch Q&A"
	c.TplName = "login.html"
}

// Signin for handling signin POST request
func (c *LoginController) Signin() {

	uname := c.Input().Get(constants.KeyUNAME)
	pwd := c.Input().Get(constants.KeyPWD)

	result, acc := biz.CheckAccount(uname, pwd)
	if result == true {

		utils.CookieInstance().Set(c.Ctx, constants.KeyUID, utils.I64toa(int64(acc.ID)), -1)
		utils.CookieInstance().Set(c.Ctx, constants.KeyUNAME, uname, -1)
		utils.CookieInstance().SetSecret(c.Ctx, constants.KeyPWD, pwd, -1)

		c.Redirect("/", 302)
		return
	}
	c.Redirect("/login/error", 302)
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
