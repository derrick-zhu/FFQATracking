package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

// IssueDetailController controller for display/edit issue in detail
type IssueDetailController struct {
	FFBaseController

	issueDetailData TIssueNewCollectionType
	issueID         int64
	currentIssue    *models.BugModel
	logHistory      []models.IssueLogModel
}

// Get handle HTTP Get request
func (c *IssueDetailController) Get() {
	c.FFBaseController.Get()

	var err error
	c.issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
	if err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}
	beego.Info(fmt.Sprintf("issue Id: %d", c.issueID))

	if c.currentIssue, err = models.BugWithID(c.issueID); err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}

	c.initVariables()
	c.initLogHistory()
	c.initPageContent()

	c.TplName = "issueDetail.html"
}

// SubmitNewLog handle POST rquest to append new issue log.
func (c *IssueDetailController) SubmitNewLog() {
	c.FFBaseController.Post()

	beego.Debug(c.Ctx.Input)
	beego.Debug(c.Input())
	beego.Debug(">>>> new comment for issue: " + c.Ctx.Input.Param(":issue"))
	beego.Debug(">>>> new comment: " + c.Input().Get("issue_comment"))

	if biz.HadLogin(c.Ctx) == false {
		c.Redirect("/login", 302)
		return
	}

	var currAcc *models.AccountModel
	var err error

	c.issueID, _ = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
	logContent := c.Input().Get("issue_comment")
	newStatus, _ := strconv.ParseInt(c.Input().Get("Status"), 10, 64)

	if c.currentIssue, err = models.BugWithID(c.issueID); err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}
	prvStatus := c.currentIssue.Status
	c.currentIssue.Status = newStatus

	if currAcc, err = biz.CurrentAccount(c.Ctx); err != nil {
		beego.Error(err)
		c.Redirect("#", 302)
		return
	}

	_, err = models.AddComment(c.issueID, int64(currAcc.ID), prvStatus, newStatus, logContent)
	if err != nil {
		beego.Error(err)
		c.Redirect("#", 302)
		return
	}

	if err = models.UpdateBug(*c.currentIssue); err != nil {
		beego.Error(err)
		c.Redirect("#", 302)
		return
	}

	c.Redirect("#", 302)
}

// UpdateIssue issue's properties changed.
func (c *IssueDetailController) UpdateIssue() {

}

// initVariables issue's properties
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

func (c *IssueDetailController) initLogHistory() {

	var err error
	c.logHistory, err = models.AllCommentsForIssue(c.issueID)
	if err != nil {
		beego.Error(err)
		return
	}

	models.SortCommentByTime(c.logHistory)

	beego.Info(c.logHistory)
}

// initPageContent initial settings in current page
func (c *IssueDetailController) initPageContent() {

	c.Data[constants.KeyIssueHTMLValue] = c.issueDetailData
	c.Data[constants.KeyIssueLogHistory] = c.logHistory
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
