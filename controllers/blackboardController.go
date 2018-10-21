package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

// BlackboardController the class of issue list
type BlackboardController struct {
	FFBaseController
}

// Get for handling issue list page HTTP GET request
func (c *BlackboardController) Get() {

	c.FFBaseController.Get()
	c.Data[constants.KeyIsBlackBoard] = 1

	var allBugs *[]models.BugModel
	var allUsers *[]models.AccountModel
	var err error

	if allBugs, err = models.AllBugsData(); err != nil {
		beego.Error(err)
	}

	if allUsers, err = models.AllAccounts(); err != nil {
		beego.Error(err)
	}

	allProjects, err := models.AllInitiatives(0, -1)
	if err != nil {
		_ = allProjects
		beego.Error(err)
		return
	}

	c.Data["allIssue"] = allBugs
	c.Data["allAccount"] = allUsers

	c.initFilterVars(allUsers, allProjects)
	c.initProjectListVar(allUsers, allProjects)
	c.initNewInitiativeVar()
	c.initNewIssueVar()

	c.TplName = "blackboardController.html"
}
