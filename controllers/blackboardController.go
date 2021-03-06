package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"strconv"

	"github.com/astaxie/beego"
)

// BlackboardController the class of issue list
type BlackboardController struct {
	FFBaseController

	allBugs     *[]models.BugModel
	allUsers    *[]models.AccountModel
	allProjects *[]models.InitiativeModel
}

// Get for handling issue list page HTTP GET request
func (c *BlackboardController) Get() {
	c.FFBaseController.Get()

	var err error
	var selSprintID int64 = -1
	var selProjID int64 = -1
	var selMSID int64 = -1

	if selSprintID, err = strconv.ParseInt(c.GetString("sprint", "-1"), 10, 64); err != nil {
		beego.Error(err)
		selSprintID = -1
	}

	if selProjID, err = strconv.ParseInt(c.GetString("proj", "-1"), 10, 64); err != nil {
		beego.Error(err)
		selProjID = -1
	}

	if selMSID, err = strconv.ParseInt(c.GetString("ms", "-1"), 10, 64); err != nil {
		beego.Error(err)
		selMSID = -1
	}

	c.commonInitForGet(selSprintID, selProjID, selMSID)
}

// private helpers

func (c *BlackboardController) commonInitForGet(selectedSprint, selectedProjID, selectedMilestoneID int64) {
	c.Data[constants.KeyIsBlackBoard] = 1

	var err error

	if c.allBugs, err = models.BugsFromProjectID(selectedSprint, selectedProjID, selectedMilestoneID, 0, -1); err != nil {
		beego.Error(err)
	}

	if c.allUsers, err = models.AllAccounts(); err != nil {
		beego.Error(err)
	}

	c.allProjects, err = c.fetchAllInitiatives()
	if err != nil {
		_ = c.allProjects
		beego.Error(err)
		return
	}

	c.Data["allIssue"] = c.allBugs
	c.Data["allAccount"] = c.allUsers
	c.Data["initiativeID"] = strconv.FormatInt(selectedProjID, 10)
	c.Data["milestoneID"] = strconv.FormatInt(selectedMilestoneID, 10)

	c.initFilterVars(selectedProjID, selectedMilestoneID, c.allUsers, c.allProjects)
	c.initProjectListVar(c.allUsers, c.allProjects)
	c.initNewInitiativeVar()
	c.initNewIssueVar()
	c.initNewMilestoneVar()

	c.TplName = "blackboardController.html"
}

func (c *BlackboardController) fetchAllInitiatives() (*[]models.InitiativeModel, error) {

	return models.AllInitiatives(0, -1)
}

func (c *BlackboardController) fetchMilestoneFilterWithInitiativeID(initID, offset, count int64) (*[]models.MilestoneModel, error) {

	result, err := models.MilestonesWithInitiative(initID, offset, count)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return result, nil
}
