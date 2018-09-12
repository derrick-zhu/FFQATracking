package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type IssueDetailController struct {
	FFBaseController

	issueDetailData TIssueNewCollectionType
	issueID         int64
	currentIssue    *models.BugModel
}

func (c *IssueDetailController) Get() {
	c.FFBaseController.Get()

	beego.Info(c.Ctx.Input)

	var err error
	c.issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
	if err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}
	beego.Info(fmt.Sprintf("issue Id: %d", c.issueID))

	c.initVariables()
	c.initPageContent()

	c.TplName = "issueDetail.html"
}

func (c *IssueDetailController) initVariables() {

	var err error
	var allUsers []models.AccountModel

	c.currentIssue, err = models.BugWithID(c.issueID)
	if err != nil {
		beego.Error(err)
		err = nil
	}

	statusData := IssueStatusData
	statusData.DefaultValue = c.currentIssue.Status

	priorityData := IssuePriorityData
	priorityData.DefaultValue = c.currentIssue.Priority

	reproductData := IssueReproductionData
	reproductData.DefaultValue = c.currentIssue.Reproductability

	//
	allUsers, err = models.AllAccounts()
	if err != nil {
		beego.Error(err)
		err = nil
	}

	allCreators := IssuePickerTemplateModel{}
	allCreators.Title = IssueCreatorKey
	allCreators.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, allCreators.Title)
	allCreators.DefaultValue = indexOf(int64(c.currentIssue.Creator), allUsers)
	allCreators.Collection = allUsers

	//
	allUsers, err = models.AllAccounts()
	if err != nil {
		beego.Error(err)
		err = nil
	}

	allAssignors := IssuePickerTemplateModel{}
	allAssignors.Title = IssueAssignorKey
	allAssignors.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, allAssignors.Title)
	allAssignors.DefaultValue = indexOf(int64(c.currentIssue.Assignor), allUsers)
	allAssignors.Collection = allUsers

	c.issueDetailData = TIssueNewCollectionType{
		IssueStatusData,
		IssuePriorityData,
		IssueReproductionData,
		allCreators,
		allAssignors,
	}
}

// initPageContent initial settings in current page
func (c *IssueDetailController) initPageContent() {

	c.Data[constants.KeyIssueHTMLValue] = c.issueDetailData
	c.Data[constants.KeyIssueData] = c.currentIssue
}

func indexOf(id int64, allAccs []models.AccountModel) int64 {

	for idx, acc := range allAccs {
		if int64(acc.ID) == id {
			return int64(idx)
		}
	}
	return -1
}
