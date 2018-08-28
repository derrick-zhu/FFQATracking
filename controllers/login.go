package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	beego.Info("??????()")
	c.Data["Title"] = "Farfetch Q&A"

	if c.Input().Get("exit") == "1" {
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {

	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")

	beego.Info(fmt.Sprintf("Post() %s, %s", uname, pwd))
	if beego.AppConfig.String("adminName") != uname ||
		beego.AppConfig.String("adminPwd") != pwd {

		c.Redirect("/login/error", 302)
		return
	}

	c.Redirect("/", 302)
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
