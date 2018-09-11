package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

// IssueListController the class of issue list
type IssueListController struct {
	FFBaseController
}

// Get for handling issue list page HTTP GET request
func (c *IssueListController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "Farfetch Q&A Tracking"
	c.Data[constants.KeyIsIssueList] = 1

	bugs, err := models.BugsWithRange(0, -1)
	if err != nil {
		beego.Error(err)
	}

	beego.Info(bugs)

	c.Data["allIssue"] = bugs

	c.TplName = "issueList.html"
}
