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

	// IssueTitleKey ...
	IssueTitleKey = "Title"
	// IssueDescriptionKey ...
	IssueDescriptionKey = "Description"
	// IssueStatusKey ...
	IssueStatusKey = "Status"
	// IssuePriorityKey ...
	IssuePriorityKey = "Priority"
	// IssueReproductionKey ...
	IssueReproductionKey = "Reproductability"
	// IssueCreatorKey ...
	IssueCreatorKey = "Creator"
	// IssueAssignorKey ...
	IssueAssignorKey = "Assignor"
)

func issuePickerKey(key string) string {
	if len(key) > 0 {
		return issueIDPrefix + key
	}
	return key
}

// TIssueNewCollectionType for issue template
type TIssueNewCollectionType []interface{}

// IssueStatusData status data (temperary)
var IssueStatusData = models.DataPickerTemplateModel{
	DataBaseTemplateModel: models.DataBaseTemplateModel{
		Title:      IssueStatusKey,
		Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssueStatusKey),
		Type:       models.Number,
	},
	DefaultValue: 0,
	Value:        0,
	Collection:   models.AllBugStatus,
}

// IssuePriorityData priority data (temperary)
var IssuePriorityData = models.DataPickerTemplateModel{
	DataBaseTemplateModel: models.DataBaseTemplateModel{
		Title:      IssuePriorityKey,
		Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssuePriorityKey),
		Type:       models.Number,
	},
	DefaultValue: 0,
	Value:        0,
	Collection:   models.AllPriorities,
}

// IssueReproductionData reproduction data (temperary)
var IssueReproductionData = models.DataPickerTemplateModel{
	DataBaseTemplateModel: models.DataBaseTemplateModel{
		Title:      IssueReproductionKey,
		Identifier: fmt.Sprintf("%s%s", issueIDPrefix, IssueReproductionKey),
		Type:       models.Number,
	},
	DefaultValue: 0,
	Value:        0,
	Collection:   models.AllReproductabilities,
}

// IssueNewController base issue create page
type IssueNewController struct {
	FFBaseController

	issueTemplateData TIssueNewCollectionType
	allCreators       models.DataPickerTemplateModel
	allAssignors      models.DataPickerTemplateModel
}

// Get for handle new issue controller GET request
func (c *IssueNewController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "New Issue"
	c.Data[constants.KeyIsIssueList] = 1

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

	c.Redirect("/issuelist", 302)
}

// NewAttchment handle append attachment POST request
func (c *IssueNewController) NewAttchment() {
	c.FFBaseController.Post()

	if _, err := helpers.SaveAttachFile(c.Ctx.Request, "myfile", constants.ServerUploadDir); err != nil {
		beego.Error(err)
	}

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

	allUsersVar := []models.VarModelProtocol{}
	for _, v := range *allUsers {
		allUsersVar = append(allUsersVar, v)
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
	c.allCreators = models.DataPickerTemplateModel{}
	c.allCreators.Title = IssueCreatorKey
	c.allCreators.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, c.allCreators.Title)
	c.allCreators.Type = models.Number
	c.allCreators.DefaultValue = int64(createorDefaultIndex)
	c.allCreators.Collection = allUsersVar

	c.issueTemplateData = append(c.issueTemplateData, c.allCreators)

	// append all assignor data into `pickData`
	c.allAssignors = models.DataPickerTemplateModel{}
	c.allAssignors.Title = IssueAssignorKey
	c.allAssignors.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, c.allAssignors.Title)
	c.allAssignors.Type = models.Number
	c.allAssignors.DefaultValue = 0
	c.allAssignors.Collection = allUsersVar

	c.issueTemplateData = append(c.issueTemplateData, c.allAssignors)
}

// initPageContent initial settings in current page
func (c *IssueNewController) initPageContent() {

	c.Data[constants.KeyIssueHTMLValue] = c.issueTemplateData
}
