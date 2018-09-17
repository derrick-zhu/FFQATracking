package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/models"
	"FFQATracking/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

// IssueDetailController controller for display/edit issue in detail
type IssueDetailController struct {
	FFBaseController
}

// Get handle HTTP Get request
func (c *IssueDetailController) Get() {
	c.FFBaseController.Get()
	beego.Debug(c)

	var err error
	var issueID int64
	var currentIssue *models.BugModel
	var issueDetailData *TIssueNewCollectionType
	var logHistory *[]models.IssueLogModel

	issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
	if err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}
	beego.Info(fmt.Sprintf("issue Id: %d", issueID))

	if currentIssue, err = models.BugWithID(issueID); err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}

	c.initVariables(&issueDetailData, currentIssue, issueID)
	c.initLogHistory(issueID, &logHistory)
	c.initPageContent(*currentIssue, *issueDetailData, *logHistory)

	c.TplName = "issueDetail.html"
}

// SubmitNewLog handle POST rquest to append new issue log.
func (c *IssueDetailController) SubmitNewLog() {
	c.FFBaseController.Post()

	for true {

		var err error
		var currentIssue *models.BugModel
		var currentAccount *models.AccountModel
		var retIssueLog *models.IssueLogModel

		nIssueID, _ := strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
		strNewComment := c.Input().Get("issue_comment")

		if currentIssue, err = models.BugWithID(nIssueID); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		if currentAccount, err = biz.CurrentAccount(c.Ctx); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		if retIssueLog, err = models.AddLogComment(currentIssue.ID, currentAccount.ID, strNewComment); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		beego.Info(retIssueLog)
		utils.MakeRedirectURL(&c.Data, 200, "#", "")
		break
	}
	c.ServeJSON()
}

// UpdateIssue issue's properties changed.
func (c *IssueDetailController) UpdateIssue() {
	c.FFBaseController.Post()

	for true {
		var err error
		var nIssueID int64
		var pIssue *models.BugModel
		// var pAccount *models.AccountModel

		if biz.HadLogin(c.Ctx) == false {
			beego.Error("login is needed.")
			utils.MakeRedirectURL(&c.Data, 302, "/login", "")
			break
		}

		nIssueID, _ = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
		// pAccount, _ = biz.CurrentAccount(c.Ctx)

		// fetch bug data
		if pIssue, err = models.BugWithID(nIssueID); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		beego.Info(pIssue)

		// get query params
		var inputMap = make(map[string]interface{})
		for k, a := range c.Input() {
			inputMap[k] = a[0]
		}

		beego.Info(inputMap)

		// update bug data property
		if err = utils.MapToStruct(inputMap, pIssue); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		// update the latest edit date
		pIssue.LastUpdateDate = utils.TimeTickSince1970()

		beego.Info(pIssue)

		if err = models.UpdateBug(pIssue); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		utils.MakeRedirectURL(&c.Data, 200, "#", "")
		break
	}

	c.ServeJSON()
}

// initVariables issue's properties
func (c *IssueDetailController) initVariables(dataSource **TIssueNewCollectionType, aIssue *models.BugModel, nIssueID int64) {

	var err error
	var allUsers *[]models.AccountModel

	aIssue, err = models.BugWithID(nIssueID)
	if err != nil {
		beego.Error(err)
		err = nil
	}

	statusData := IssueStatusData
	statusData.DefaultValue = aIssue.Status
	statusData.ID = nIssueID

	priorityData := IssuePriorityData
	priorityData.DefaultValue = aIssue.Priority
	priorityData.ID = nIssueID

	reproductData := IssueReproductionData
	reproductData.DefaultValue = aIssue.Reproductability
	reproductData.ID = nIssueID

	//
	allUsers, err = models.AllAccounts()
	if err != nil {
		beego.Error(err)
		err = nil
	}

	allCreators := IssuePickerTemplateModel{}
	allCreators.Title = IssueCreatorKey
	allCreators.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, allCreators.Title)
	allCreators.DefaultValue = indexOf(int64(aIssue.Creator), *allUsers)
	allCreators.Collection = *allUsers
	allCreators.ID = nIssueID

	//
	allUsers, err = models.AllAccounts()
	if err != nil {
		beego.Error(err)
		err = nil
	}

	allAssignors := IssuePickerTemplateModel{}
	allAssignors.Title = IssueAssignorKey
	allAssignors.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, allAssignors.Title)
	allAssignors.DefaultValue = indexOf(int64(aIssue.Assignor), *allUsers)
	allAssignors.Collection = *allUsers
	allAssignors.ID = nIssueID

	*dataSource = &TIssueNewCollectionType{
		statusData,
		priorityData,
		reproductData,
		allCreators,
		allAssignors,
	}
}

func (c *IssueDetailController) initLogHistory(nIssueID int64, logHistory **[]models.IssueLogModel) {

	var err error
	var logs *[]models.IssueLogModel

	if logs, err = models.AllCommentsForIssue(nIssueID); err != nil {
		beego.Error(err)
		return
	}

	models.SortCommentByTime(logs)

	*logHistory = logs

	beego.Debug(logHistory)
}

// initPageContent initial settings in current page
func (c *IssueDetailController) initPageContent(aIssue models.BugModel, dataSource TIssueNewCollectionType, logHistory []models.IssueLogModel) {

	c.Data[constants.KeyIssueHTMLValue] = dataSource
	c.Data[constants.KeyIssueLogHistory] = logHistory
	c.Data[constants.KeyIssueData] = aIssue
}

func indexOf(id int64, allAccs []models.AccountModel) int64 {

	for idx, acc := range allAccs {
		if int64(acc.ID) == id {
			return int64(idx)
		}
	}
	return -1
}
