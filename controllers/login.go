package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	c.Data["Title"] = "Farfetch Q&A"
	c.TplName = "login.html"
}

func (c *LoginController) Post() {

	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")

	if beego.AppConfig.String("adminName") != uname ||
		beego.AppConfig.String("adminPwd") != pwd {

		c.Redirect("/login/error", 302)
		return
	}

	c.Redirect("/", 302)
	return
}

func (c *LoginController) Error() {

	beego.Error("something wrong")
	c.Redirect("/login", 302)
}
