package controllers

import (
	"FFQATracking/constants"
	"fmt"

	"github.com/astaxie/beego"
)

// NewIssueController base issue create page
type NewIssueController struct {
	FFBaseController
}

// Get for handle new issue controller GET request
func (c *NewIssueController) Get() {
	c.FFBaseController.Get()
	c.Data[constants.Title] = "New Issue"
	c.Data[constants.KeyIsIssueList] = 1

	c.TplName = "newIssue.html"
}

// Post for handle new issue controller POST request
func (c *NewIssueController) Post() {
	c.FFBaseController.Post()

	bugTitle := c.Input().Get("issueTitle")
	issueContent := c.Input().Get("issueContent")

	beego.Info(fmt.Sprintf("title: %s, description: %s", bugTitle, issueContent))

	c.Redirect("/issuelist", 302)
}
