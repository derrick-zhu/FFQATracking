package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"FFQATracking/utils"
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
)

// IssueDetailController controller for display/edit issue in detail
type IssueDetailController struct {
	FFBaseController

	// issueDetailData TIssueNewCollectionType
	// issueID         int64
	// currentIssue    *models.BugModel
	// logHistory      []models.IssueLogModel
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
	beego.Debug(c)

	beego.Debug(c.Ctx.Input)
	beego.Debug(c.Input())
	beego.Debug(">>>> new comment for issue: " + c.Ctx.Input.Param(":issue"))
	beego.Debug(">>>> new comment: " + c.Input().Get("issue_comment"))

	for true {
		nIssueID, _ := strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
		param := utils.CovertRequestInputToMap(c.Input())
		beego.Error(param)

		var err error
		var currentIssue *models.BugModel

		if currentIssue, err = models.BugWithID(nIssueID); err != nil {
			beego.Error(err)
			return
		}

		beego.Error(currentIssue)

		newIssue := models.BugModel{}

		err = utils.MapToStruct(param, &newIssue)
		if err != nil {
			beego.Error(err)
		}

		beego.Error(newIssue)

		bSame := reflect.DeepEqual(currentIssue, newIssue)
		if bSame == false {
		}

		// var nIssueID int64
		// var currentIssue *models.BugModel

		// for true {
		// 	if biz.HadLogin(c.Ctx) == false {
		// 		utils.MakeRedirectURL(&c.Data, 302, "/login", "")
		// 		break
		// 	}

		// 	var currAcc *models.AccountModel
		// 	var err error

		// 	nIssueID, _ = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
		// 	logContent := c.Input().Get("issue_comment")
		// 	newStatus, _ := strconv.ParseInt(c.Input().Get("Status"), 10, 64)

		// 	if currentIssue, err = models.BugWithID(nIssueID); err != nil {
		// 		beego.Error(err)
		// 		utils.MakeRedirectURL(&c.Data, 302, "/issuelist", "")
		// 		break
		// 	}
		// 	prvStatus := currentIssue.Status
		// 	currentIssue.Status = newStatus

		// 	if currAcc, err = biz.CurrentAccount(c.Ctx); err != nil {
		// 		beego.Error(err)
		// 		utils.MakeRedirectURL(&c.Data, 302, "#", "")
		// 		break
		// 	}

		// 	_, err = models.AddComment(nIssueID, int64(currAcc.ID), prvStatus, newStatus, logContent)
		// 	if err != nil {
		// 		beego.Error(err)
		// 		utils.MakeRedirectURL(&c.Data, 302, "#", "")
		// 		break
		// 	}

		// 	if err = models.UpdateBug(*currentIssue); err != nil {
		// 		beego.Error(err)
		// 		utils.MakeRedirectURL(&c.Data, 302, "#", "")
		// 		break
		// 	}

		// 	utils.MakeRedirectURL(&c.Data, 302, "#", "")
		// 	break
		// }

		utils.MakeRedirectURL(&c.Data, 302, "#", "")
		break
	}
	c.ServeJSON()
}

// UpdateIssue issue's properties changed.
func (c *IssueDetailController) UpdateIssue() {
	c.FFBaseController.Post()

	beego.Debug(">>>>> UpdateIssue ->")
	beego.Info(c.Ctx.Input)
	beego.Debug(c.Input())

	nIssueID, _ := strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)

	var inputMap = make(map[string]interface{})
	for k, a := range c.Input() {
		inputMap[k] = a[0]
	}

	var pIssue = &models.BugModel{}
	pIssue.ID = nIssueID
	utils.MapToStruct(inputMap, pIssue)

	beego.Error(inputMap)
	beego.Error(pIssue)

	for true {
		// nIssueID := c.Ctx.Input.Param(":issue")
		break
	}

	utils.MakeRedirectURL(&c.Data, 302, "#", "")
	c.ServeJSON()
}

// initVariables issue's properties
func (c *IssueDetailController) initVariables(dataSource **TIssueNewCollectionType, aIssue *models.BugModel, nIssueID int64) {

	var err error
	var allUsers []models.AccountModel

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
	allCreators.DefaultValue = indexOf(int64(aIssue.Creator), allUsers)
	allCreators.Collection = allUsers
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
	allAssignors.DefaultValue = indexOf(int64(aIssue.Assignor), allUsers)
	allAssignors.Collection = allUsers
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
	logs, err = models.AllCommentsForIssue(nIssueID)
	if err != nil {
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
