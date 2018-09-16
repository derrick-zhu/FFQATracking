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
	c.Data[constants.KeyIsIssueList] = 1

	var allBugs *[]models.BugModel
	var allUsers *[]models.AccountModel
	var err error

	if allBugs, err = models.AllBugsData(); err != nil {
		beego.Error(err)
	}

	if allUsers, err = models.AllAccounts(); err != nil {
		beego.Error(err)
	}

	c.Data["allIssue"] = allBugs
	c.Data["allAccount"] = allUsers

	c.TplName = "issueList.html"
}
