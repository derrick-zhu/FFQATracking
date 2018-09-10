package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	issueIDPrefix = "picker-"

	IssueTitleKey        = "Title"
	IssueDescriptionKey  = "Description"
	IssueStatusKey       = "Status"
	IssuePriorityKey     = "Priority"
	IssueReproductionKey = "Reproduction"
	IssueCreatorKey      = "Creator"
	IssueAssignorKey     = "Assignor"
)

func PickerKey(key string) string {
	if len(key) > 0 {
		return issueIDPrefix + key
	}
	return key
}

// IssuePickerTemplateModel class template
type IssuePickerTemplateModel struct {
	Title         string
	Identifier    string
	DefaultValue  int64
	Value         int64
	Collection    []string
	ExtCollection []int64
}

// TIssueNewCollectionType for issue template
type TIssueNewCollectionType []IssuePickerTemplateModel

// IssueStatusData status data (temperary)
var IssueStatusData = IssuePickerTemplateModel{
	Title:        IssueStatusKey,
	Identifier:   fmt.Sprintf("%s%s", issueIDPrefix, IssueStatusKey),
	DefaultValue: 0,
	Value:        0,
	Collection:   models.EnumAllBugsStatus(),
}

// IssuePriorityData priority data (temperary)
var IssuePriorityData = IssuePickerTemplateModel{
	Title:        IssuePriorityKey,
	Identifier:   fmt.Sprintf("%s%s", issueIDPrefix, IssuePriorityKey),
	DefaultValue: 0,
	Value:        0,
	Collection:   models.EnumAllBugsPriority(),
}

// IssueReproductionData reproduction data (temperary)
var IssueReproductionData = IssuePickerTemplateModel{
	Title:        IssueReproductionKey,
	Identifier:   fmt.Sprintf("%s%s", issueIDPrefix, IssueReproductionKey),
	DefaultValue: 0,
	Value:        0,
	Collection:   []string{"100%", "80%", "60%", "40%", "20%"},
}

// IssueController base issue create page
type IssueController struct {
	FFBaseController

	issueTemplateData TIssueNewCollectionType
	allCreators       IssuePickerTemplateModel
	allAssignors      IssuePickerTemplateModel
}

// Get for handle new issue controller GET request
func (c *IssueController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "New Issue"
	c.Data[constants.KeyIsIssueList] = 1

	c.initPageVariables()
	c.initPageContent()

	c.TplName = "issue.html"
}

// Post for handle new issue controller POST request
func (c *IssueController) Post() {
	c.FFBaseController.Post()

	bugTitle := c.Input().Get("issueTitle")
	issueContent := c.Input().Get("issueContent")

	beego.Info(fmt.Sprintf("title: %s, description: %s", bugTitle, issueContent))

	c.Redirect("/issuelist", 302)
}

// Create the method for creating issue
func (c *IssueController) Create() {
	// beego.Info(c.Input())

	strStatus := c.Input().Get(PickerKey(IssueStatusKey))
	status := models.BugStatusWithString(strStatus)

	strPriority := c.Input().Get(PickerKey(IssuePriorityKey))
	priority := models.BugPriorityWithString(strPriority)

	strReproduct := c.Input().Get(PickerKey(IssueReproductionKey))
	reproduct := models.BugReproductabilityWithString(strReproduct)

	beego.Info(fmt.Sprintf("status: %s -> %d", strStatus, status))
	beego.Info(fmt.Sprintf("priority: %s -> %d", strPriority, priority))
	beego.Info(fmt.Sprintf("reproduct: %s -> %d", strReproduct, reproduct))

	c.Redirect("#", 302)
}

// MARK - private helpers

func (c *IssueController) initPageVariables() {
	// fetch all user data
	allUsers, err := models.AllAccounts()
	if err != nil {
		beego.Error(err)
	}

	c.issueTemplateData = TIssueNewCollectionType{

		IssueStatusData, IssuePriorityData, IssueReproductionData,
	}

	// append all creators data into `pickData`
	c.allCreators = IssuePickerTemplateModel{}
	c.allCreators.Title = IssueCreatorKey
	c.allCreators.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, c.allCreators.Title)
	c.allCreators.DefaultValue = 0

	for _, eachUser := range allUsers {
		c.allCreators.Collection = append(c.allCreators.Collection, eachUser.Name)
		c.allCreators.ExtCollection = append(c.allCreators.ExtCollection, int64(eachUser.ID))
	}
	c.issueTemplateData = append(c.issueTemplateData, c.allCreators)

	// append all assignor data into `pickData`
	c.allAssignors = IssuePickerTemplateModel{}
	c.allAssignors.Title = IssueAssignorKey
	c.allAssignors.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, c.allAssignors.Title)
	c.allAssignors.DefaultValue = 0
	for _, eachAssignor := range allUsers {
		c.allAssignors.Collection = append(c.allAssignors.Collection, eachAssignor.Name)
		c.allAssignors.ExtCollection = append(c.allAssignors.ExtCollection, int64(eachAssignor.ID))
	}
	c.issueTemplateData = append(c.issueTemplateData, c.allAssignors)
}

// initPageContent initial settings in current page
func (c *IssueController) initPageContent() {

	c.Data[constants.KeyIssueHTMLValue] = c.issueTemplateData
}
