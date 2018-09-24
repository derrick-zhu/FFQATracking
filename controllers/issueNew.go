package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/constants"
	"FFQATracking/helpers"
	"FFQATracking/models"
	"FFQATracking/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	issueIDPrefix = ""

	IssueTitleKey        = "Title"
	IssueDescriptionKey  = "Description"
	IssueStatusKey       = "Status"
	IssuePriorityKey     = "Priority"
	IssueReproductionKey = "Reproductability"
	IssueCreatorKey      = "Creator"
	IssueAssignorKey     = "Assignor"
)

func issuePickerKey(key string) string {
	if len(key) > 0 {
		return issueIDPrefix + key
	}
	return key
}

// IssuePickerTemplateModel class template
type IssuePickerTemplateModel struct {
	ID           int64 // just ID for any index num
	Title        string
	Identifier   string
	DefaultValue int64
	Value        int64
	Collection   interface{}
}

// TIssueNewCollectionType for issue template
type TIssueNewCollectionType []IssuePickerTemplateModel

// TIssueAttachmentType for issue's attachment
type TIssueAttachmentType []models.AttachmentModel

// IssueStatusData status data (temperary)
var IssueStatusData = IssuePickerTemplateModel{
	Title:        IssueStatusKey,
	Identifier:   fmt.Sprintf("%s%s", issueIDPrefix, IssueStatusKey),
	DefaultValue: 0,
	Value:        0,
	Collection:   models.AllBugStatus,
}

// IssuePriorityData priority data (temperary)
var IssuePriorityData = IssuePickerTemplateModel{
	Title:        IssuePriorityKey,
	Identifier:   fmt.Sprintf("%s%s", issueIDPrefix, IssuePriorityKey),
	DefaultValue: 0,
	Value:        0,
	Collection:   models.AllPriorities,
}

// IssueReproductionData reproduction data (temperary)
var IssueReproductionData = IssuePickerTemplateModel{
	Title:        IssueReproductionKey,
	Identifier:   fmt.Sprintf("%s%s", issueIDPrefix, IssueReproductionKey),
	DefaultValue: 0,
	Value:        0,
	Collection:   models.AllReproductabilities,
}

// IssueNewController base issue create page
type IssueNewController struct {
	FFBaseController

	issueTemplateData TIssueNewCollectionType
	issueAttachData   TIssueAttachmentType
	allCreators       IssuePickerTemplateModel
	allAssignors      IssuePickerTemplateModel
}

// just for store the attach session key
var gAttachSessionKeyForNewIssue int64

// Get for handle new issue controller GET request
func (c *IssueNewController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "New Issue"
	c.Data[constants.KeyIsIssueList] = 1

	if gAttachSessionKeyForNewIssue != 0 {
		biz.SharedAttachManager().RemoveSession(gAttachSessionKeyForNewIssue)
	}
	gAttachSessionKeyForNewIssue = biz.SharedAttachManager().NewSession()

	c.initPageVariables()
	c.initPageContent()

	c.TplName = "issueNew.html"
}

// Create the method for creating issue
func (c *IssueNewController) Create() {

	var err error
	var nStatus int64
	var nPriority int64
	var nReproduct int64
	var nCreatorID int64
	var nAssignorID int64

	title := c.Input().Get(issuePickerKey(IssueTitleKey))
	description := c.Input().Get(issuePickerKey(IssueDescriptionKey))
	status := c.Input().Get(issuePickerKey(IssueStatusKey))
	priority := c.Input().Get(issuePickerKey(IssuePriorityKey))
	reproduct := c.Input().Get(issuePickerKey(IssueReproductionKey))
	creatorID := c.Input().Get(issuePickerKey(IssueCreatorKey))
	assignorID := c.Input().Get(issuePickerKey(IssueAssignorKey))

	// beego.Debug(fmt.Sprintf("title: %s", title))
	// beego.Debug(fmt.Sprintf("description: %s", description))
	// beego.Debug(fmt.Sprintf("status: %s -> %s", IssueStatusKey, status))
	// beego.Debug(fmt.Sprintf("priority: %s -> %s", IssuePriorityKey, priority))
	// beego.Debug(fmt.Sprintf("reproduct: %s -> %s", IssueReproductionKey, reproduct))
	// beego.Debug(fmt.Sprintf("creator: %s -> %s", IssueCreatorKey, creatorID))
	// beego.Debug(fmt.Sprintf("assignor: %s -> %s", IssueAssignorKey, assignorID))

	nStatus, err = strconv.ParseInt(status, 10, 64)
	nPriority, err = strconv.ParseInt(priority, 10, 64)
	nReproduct, err = strconv.ParseInt(reproduct, 10, 64)
	nCreatorID, err = strconv.ParseInt(creatorID, 10, 64)
	nAssignorID, err = strconv.ParseInt(assignorID, 10, 64)

	_, err = models.AddBug(title, description, nStatus, nPriority, nCreatorID, nAssignorID, nReproduct)
	if err != nil {
		beego.Error(err)
		c.Redirect("#", 302)

		return
	}

	// 提交完issue之后，需要清理持有的attachment session
	if gAttachSessionKeyForNewIssue != 0 {
		biz.SharedAttachManager().RemoveSession(gAttachSessionKeyForNewIssue)
		gAttachSessionKeyForNewIssue = 0
	}

	c.Redirect("/issuelist", 302)
}

// NewAttchment handle append attachment POST request
func (c *IssueNewController) NewAttchment() {
	c.FFBaseController.Post()

	attachName, err := helpers.SaveAttachFile(c.Ctx.Request, "myfile", constants.ServerUploadDir)
	if err != nil {
		beego.Error(err)
	}

	// TODO: add attachment into db here.
	attachSession := biz.SharedAttachManager().SessionWithID(gAttachSessionKeyForNewIssue)
	attachSession.AppendAttachement(attachName)

	utils.MakeRedirectURL(&c.Data, 302, "#", "")

	c.ServeJSON()
}

// MARK - private helpers

func (c *IssueNewController) initPageVariables() {
	// fetch all user data
	allUsers, err := models.AllAccounts()
	if err != nil {
		beego.Error(err)
	}

	c.issueTemplateData = TIssueNewCollectionType{

		IssueStatusData, IssuePriorityData, IssueReproductionData,
	}

	acc, err := biz.CurrentAccount(c.Ctx)
	createorDefaultIndex := 0 // the index of current logged in user in allUser array

	// cause of making current logged in user as a default creator accout, find and save the index of current accout in the allUser array
	if acc != nil {
		for index, eachAcc := range *allUsers {
			if eachAcc.ID == acc.ID {
				createorDefaultIndex = index
			}
		}
	}

	// append all creators data into `pickData`
	c.allCreators = IssuePickerTemplateModel{}
	c.allCreators.Title = IssueCreatorKey
	c.allCreators.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, c.allCreators.Title)
	c.allCreators.DefaultValue = int64(createorDefaultIndex)
	c.allCreators.Collection = *allUsers

	c.issueTemplateData = append(c.issueTemplateData, c.allCreators)

	// append all assignor data into `pickData`
	c.allAssignors = IssuePickerTemplateModel{}
	c.allAssignors.Title = IssueAssignorKey
	c.allAssignors.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, c.allAssignors.Title)
	c.allAssignors.DefaultValue = 0
	c.allAssignors.Collection = *allUsers

	c.issueTemplateData = append(c.issueTemplateData, c.allAssignors)
}

// initPageContent initial settings in current page
func (c *IssueNewController) initPageContent() {

	c.Data[constants.KeyIssueHTMLValue] = c.issueTemplateData
}
