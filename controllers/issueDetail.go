package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/helpers"
	"FFQATracking/models"
	"FFQATracking/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

const (
	// KDraftCommentKey key for temp saving issue comment
	KDraftCommentKey string = "issue-detail-draft-comment"
)

// IssueDetailLogModel issue log wrapped model
type IssueDetailLogModel struct {
	models.IssueLogModel

	CreatorName      string
	CreatorAvatar    string
	TimeDisplay      string
	PrvStatusStr     string
	NewStatusStr     string
	IsViewersComment bool
}

func (c *IssueDetailLogModel) initWith(other *models.IssueLogModel, acc *models.AccountModel, isViewersLog bool) {
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
	c.IsViewersComment = isViewersLog
	c.TimeDisplay = utils.StandardFormatedTimeFromTick(other.Time)
}

// IssueDetailController controller for display/edit issue in detail
type IssueDetailController struct {
	FFBaseController
}

// Get handle HTTP Get request
func (c *IssueDetailController) Get() {
	c.FFBaseController.Get()

	c.setupNormalResponseData()
	c.TplName = "issueDetail.html"
}

// SubmitNewLog handle POST rquest to append new issue log.
func (c *IssueDetailController) SubmitNewLog() {
	c.FFBaseController.Post()

	for {

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

	for {
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
		c.Redirect("/blackboard", 303)
		return
	}

	fp, err := helpers.SaveAttachFile(c.Ctx.Request, "attachImage", "static/upload/")
	if err != nil {
		beego.Error(err)
	}

	curComment := c.GetString("issue_comment")
	if strings.HasPrefix(fp, "/") == false {
		fp = "/" + fp
	}

	newComment := curComment + "![" + fp + "](" + fp + ")"
	utils.CookieInstance().Set(c.Ctx, KDraftCommentKey, newComment, 1<<32-1)

	finalURL := fmt.Sprintf("/issuedetail/%d/", issueID)
	utils.MakeRedirectURLWithUserInfo(&c.Data, 303, finalURL, "", fp)

	// c.Redirect(finalURL, 303)
	c.ServeJSON()
}

// DeleteComment delete current user's comment in issue's log history
func (c *IssueDetailController) DeleteComment() {
	c.FFBaseController.Post()

	beego.Info(c.Input())

	var issueID int64
	var issueCommentID int64
	var err error

	for {

		var finalURL = fmt.Sprintf("/issuedetail/%d", issueID)

		if issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64); err != nil {
			beego.Error(err)
			err = nil
			c.Redirect("#", 404)
			break
		}

		if issueCommentID, err = strconv.ParseInt(c.Input().Get("comment"), 10, 64); err != nil {
			beego.Error(err)
			err = nil
			c.Redirect("#", 404)
			break
		}

		if err = models.RemoveComment(issueID, issueCommentID); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, finalURL, "fails in remove issue log")
		}

		utils.MakeRedirectURL(&c.Data, 200, finalURL, "")
		break
	}

	c.ServeJSON()
}

// initVariables issue's properties
func (c *IssueDetailController) initVariables(dataSource **TIssueNewCollectionType, aIssue *models.BugModel, nIssueID int64, allUsers []models.VarModelProtocol) {

	var err error

	aIssue, err = models.BugWithID(nIssueID)
	if err != nil {
		beego.Error(err)
		err = nil
	}

	*dataSource = &TIssueNewCollectionType{
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				ID:         nIssueID,
				Title:      IssueStatusKey,
				Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssueStatusKey),
				Type:       models.Number,
			},
			DefaultValue: aIssue.Status,
			Value:        aIssue.Status,
			Collection:   models.AllBugStatus,
		},
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				ID:         nIssueID,
				Title:      IssuePriorityKey,
				Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssuePriorityKey),
				Type:       models.Number,
			},
			DefaultValue: aIssue.Priority,
			Value:        aIssue.Priority,
			Collection:   models.AllPriorities,
		},
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				ID:         nIssueID,
				Title:      IssueReproductionKey,
				Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssueReproductionKey),
				Type:       models.Number,
			},
			DefaultValue: aIssue.Reproductability,
			Value:        aIssue.Reproductability,
			Collection:   models.AllReproductabilities,
		},
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				ID:         nIssueID,
				Title:      IssueCreatorKey,
				Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssueCreatorKey),
				Type:       models.Number,
			},
			DefaultValue: indexOf(int64(aIssue.Creator), allUsers),
			Value:        indexOf(int64(aIssue.Creator), allUsers),
			Collection:   allUsers,
		},
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				ID:         nIssueID,
				Title:      IssueAssignorKey,
				Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssueAssignorKey),
				Type:       models.Number,
			},
			DefaultValue: indexOf(int64(aIssue.Assignor), allUsers),
			Value:        indexOf(int64(aIssue.Assignor), allUsers),
			Collection:   allUsers,
		},
	}
}

// initLogHistory issue's comment history (log history)
func (c *IssueDetailController) initLogHistory(nIssueID int64, logHistory **[]IssueDetailLogModel, allUsers *[]models.AccountModel) {

	var err error
	var viewerAcc *models.AccountModel
	var logs *[]models.IssueLogModel

	if logs, err = models.AllCommentsForIssue(nIssueID); err != nil {
		beego.Error(err)
		return
	}

	if viewerAcc, err = biz.AccountManagerInstance().CurrentAccount(c.Ctx); err != nil {
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

		isViewersLog := viewerAcc.ID == eachLog.CreatorID

		newIssueLog := IssueDetailLogModel{}
		newIssueLog.initWith(&eachLog, pAcc, isViewersLog)

		issueLogs = append(issueLogs, newIssueLog)
	}

	*logHistory = &issueLogs
}

func (c *IssueDetailController) initCommentContent() {
	draftComment := utils.CookieInstance().Get(c.Ctx, KDraftCommentKey)
	if len(draftComment) == 0 {
		return
	}
	c.Data[KDraftCommentKey] = draftComment
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

func indexOf(id int64, allAccs []models.VarModelProtocol) int64 {

	for idx, acc := range allAccs {
		if int64(acc.Type()) == id {
			return int64(idx)
		}
	}
	return -1
}

/**
funcs for golang template
*/

func (c *IssueDetailController) setupNormalResponseData() {
	var err error
	var issueID int64
	var currentIssue *models.BugModel
	var issueDetailData *TIssueNewCollectionType
	var logHistory *[]IssueDetailLogModel
	var allUsers *[]models.AccountModel

	if issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64); err != nil {
		beego.Error(err)
		c.Redirect("/blackboard", 302)
		return
	}

	if currentIssue, err = models.BugWithID(issueID); err != nil {
		beego.Error(err)
		c.Redirect("/blackboard", 302)
		return
	}

	if allUsers, err = models.AllAccounts(); err != nil {
		beego.Error(err)
		err = nil
	}

	allUserVars := []models.VarModelProtocol{}
	for _, v := range *allUsers {
		allUserVars = append(allUserVars, v)
	}

	c.initVariables(&issueDetailData, currentIssue, issueID, allUserVars)
	c.initLogHistory(issueID, &logHistory, allUsers)
	c.initPageContent(*currentIssue, *issueDetailData, *logHistory)
}
