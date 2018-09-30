package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/helpers"
	"FFQATracking/models"
	"FFQATracking/utils"
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
)

// IssueDetailLogModel issue log wrapped model
type IssueDetailLogModel struct {
	models.IssueLogModel

	CreatorName   string
	CreatorAvatar string
	TimeDisplay   string
	PrvStatusStr  string
	NewStatusStr  string
}

func (c *IssueDetailLogModel) initWith(other *models.IssueLogModel, acc *models.AccountModel) {
	c.IssueLogModel.ID = other.ID
	c.IssueLogModel.IssueID = other.IssueID
	c.IssueLogModel.Type = other.Type
	c.IssueLogModel.Content = other.Content

	c.IssueLogModel.CreatorID = other.CreatorID
	c.IssueLogModel.StatusTitle = other.StatusTitle
	c.IssueLogModel.PrvStatus = other.PrvStatus
	c.IssueLogModel.NewStatus = other.NewStatus

	c.CreatorName = acc.Name
	c.CreatorAvatar = acc.Avatar
	c.TimeDisplay = utils.StandardFormatedTimeFromTick(other.Time)
}

// IssueDetailController controller for display/edit issue in detail
type IssueDetailController struct {
	FFBaseController
}

// Get handle HTTP Get request
func (c *IssueDetailController) Get() {
	c.FFBaseController.Get()

	var err error
	var issueID int64
	var currentIssue *models.BugModel
	var issueDetailData *TIssueNewCollectionType
	var logHistory *[]IssueDetailLogModel
	var allUsers *[]models.AccountModel

	if issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64); err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}

	if currentIssue, err = models.BugWithID(issueID); err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}

	if allUsers, err = models.AllAccounts(); err != nil {
		beego.Error(err)
		err = nil
	}

	c.initVariables(&issueDetailData, currentIssue, issueID, allUsers)
	c.initLogHistory(issueID, &logHistory, allUsers)
	c.initPageContent(*currentIssue, *issueDetailData, *logHistory)

	c.TplName = "issueDetail.html"
}

// Post handle the POST request (nothing to do with it)
func (c *IssueDetailController) Post() {
	c.FFBaseController.Post()

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
		var pAccount *models.AccountModel

		var oldStatus, newStatus int64

		if biz.HadLogin(c.Ctx) == false {
			beego.Error("login is needed.")
			utils.MakeRedirectURL(&c.Data, 302, "/login", "")
			break
		}

		nIssueID, _ = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
		pAccount, _ = biz.CurrentAccount(c.Ctx)

		// fetch bug data
		if pIssue, err = models.BugWithID(nIssueID); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}

		beego.Info(pIssue)

		// get query params
		var inputMap = make(map[string]interface{})
		var kk string
		var vv []string

		for kk, vv = range c.Input() {
			inputMap[kk] = vv[0]
		}
		beego.Info(inputMap)

		// update bug data property
		oldStatus = models.IntValueInIssue(kk, pIssue)
		if err = utils.MapToStruct(inputMap, pIssue); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 302, "#", "")
			break
		}
		newStatus = models.IntValueInIssue(kk, pIssue)

		// save the status log into db
		if oldStatus != newStatus {
			// var log *models.IssueLogModel
			if _, err = models.AddLogStatus(nIssueID, pAccount.ID, kk, oldStatus, newStatus); err != nil {
				beego.Error(err)
				utils.MakeRedirectURL(&c.Data, 302, "#", "")
				break
			}
			// beego.Error(log)
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

// NewAttachment new attachment request (POST)
func (c *IssueDetailController) NewAttachment() {
	c.FFBaseController.Post()

	var issueID int64
	var err error
	if issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64); err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}

	fp, err := helpers.SaveAttachFile(c.Ctx.Request, "attachImage", "static/upload/")
	if err != nil {
		beego.Error(err)
	}

	finalURL := fmt.Sprintf("/issuedetail/%d/#", issueID)
	utils.MakeRedirectURLWithUserInfo(&c.Data, 303, finalURL, "", fp)

	c.ServeJSON()
}

// initVariables issue's properties
func (c *IssueDetailController) initVariables(dataSource **TIssueNewCollectionType, aIssue *models.BugModel, nIssueID int64, allUsers *[]models.AccountModel) {

	var err error

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

	// all account
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

func (c *IssueDetailController) initLogHistory(nIssueID int64, logHistory **[]IssueDetailLogModel, allUsers *[]models.AccountModel) {

	var err error
	var logs *[]models.IssueLogModel

	if logs, err = models.AllCommentsForIssue(nIssueID); err != nil {
		beego.Error(err)
		return
	}

	models.SortCommentByTime(logs)

	// Q: why create new array to carry the issue log data?? why?
	// A: I was confuse about GO's memory managment
	var issueLogs = []IssueDetailLogModel{}
	for _, eachLog := range *logs {

		var pAcc *models.AccountModel
		for _, eachAcc := range *allUsers {
			if eachAcc.ID == eachLog.CreatorID {
				pAcc = &eachAcc
				break
			}
		}

		newIssueLog := IssueDetailLogModel{}
		newIssueLog.initWith(&eachLog, pAcc)

		issueLogs = append(issueLogs, newIssueLog)
	}

	*logHistory = &issueLogs
}

// initPageContent initial settings in current page
func (c *IssueDetailController) initPageContent(aIssue models.BugModel, dataSource TIssueNewCollectionType, logHistory []IssueDetailLogModel) {

	c.Data[constants.KeyIssueHTMLValue] = dataSource
	c.Data[constants.KeyIssueLogHistory] = logHistory
	c.Data[constants.KeyIssueData] = aIssue
}

/**
public functions in issue detail page
*/

func indexOf(id int64, allAccs []models.AccountModel) int64 {

	for idx, acc := range allAccs {
		if int64(acc.ID) == id {
			return int64(idx)
		}
	}
	return -1
}

/**
funcs for golang template
*/

func init() {
	beego.AddFuncMap("isLastItemIn", isLastItemIn)
}

func isLastItemIn(x int, a interface{}) bool {
	return (x == (reflect.ValueOf(a).Len() - 1))
}
