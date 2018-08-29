package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/utils"
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	beego.Info("??????()")
	op := c.Input().Get("op")
	beego.Info(fmt.Sprintf("GET op = %s", op))

	c.Data["Title"] = "Farfetch Q&A"

	if c.Input().Get("exit") == "1" {
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Signin() {

	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")

	if biz.CheckAccount(uname, pwd) {

		utils.CookieInstance().Set(c.Ctx, "uname", uname, -1)
		utils.CookieInstance().SetSecret(c.Ctx, "pwd", pwd, -1)

		c.Redirect("/", 302)
		return
	}
	c.Redirect("/login/error", 302)
}

func (this *LoginController) Signup() {
	this.Redirect("/signup", 302)
	return
}

// func (c *LoginController) Signup() {

// 	uname := c.Input().Get("uname")
// 	pwd := c.Input().Get("pwd")
// 	beego.Info(fmt.Sprintf("Signup() %s, %s", uname, pwd))
// 	c.Redirect("/", 302)
// 	return
// }

// func (c *LoginController) Error() {

// 	beego.Error("something wrong")
// 	c.Redirect("/login", 302)
// }
