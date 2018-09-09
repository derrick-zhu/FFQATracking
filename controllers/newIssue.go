package controllers

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	IssueTitleKey        = "Title:"
	IssueDescriptionKey  = "Description:"
	IssueStatusKey       = "Status:"
	IssuePriorityKey     = "Priority:"
	IssueReproductionKey = "Reproduction:"
	IssueCreatorKey      = "Creator:"
	IssueAssignorKey     = "Assignor:"
)

// IssuePickerTemplateModel class template
type IssuePickerTemplateModel struct {
	Title        string
	DefaultValue int
	Collection   []string
}

// TIssueNewCollectionType for issue template
type TIssueNewCollectionType map[string]IssuePickerTemplateModel

// IssueStatusData status data (temperary)
var IssueStatusData = IssuePickerTemplateModel{
	Title:        IssueStatusKey,
	DefaultValue: 0,
	Collection:   []string{"New", "Fixed", "Reopen", "Confirm", "Close", "Not a bug", "Will not fix", "Delay", "Must be fix"},
}

// IssuePriorityData priority data (temperary)
var IssuePriorityData = IssuePickerTemplateModel{
	Title:        IssuePriorityKey,
	DefaultValue: 0,
	Collection:   []string{"Urgent", "Important", "High", "Middle", "Low", "Question", "Suggestion"},
}

// IssueReproductionData reproduction data (temperary)
var IssueReproductionData = IssuePickerTemplateModel{
	Title:        IssueReproductionKey,
	DefaultValue: 0,
	Collection:   []string{"100%", "80%", "60%", "40%", "20%"},
}

// NewIssueController base issue create page
type NewIssueController struct {
	FFBaseController

	issueTemplateData TIssueNewCollectionType
}

// Get for handle new issue controller GET request
func (c *NewIssueController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "New Issue"
	c.Data[constants.KeyIsIssueList] = 1

	c.initPageVariables()
	c.initPageContent()

	c.TplName = "newIssue.html"
}

// Post for handle new issue controller POST request
func (c *NewIssueController) Post() {
	c.FFBaseController.Post()

	bugTitle := c.Input().Get("issueTitle")
	issueContent := c.Input().Get("issueContent")

	beego.Info(fmt.Sprintf("title: %s, description: %s", bugTitle, issueContent))

	c.Redirect("/issuelist", 302)
}

func (c *NewIssueController) initPageVariables() {
	// fetch all user data
	allUsers, err := models.AllAccounts()
	if err != nil {
		beego.Error(err)
	}

	c.issueTemplateData = TIssueNewCollectionType{

		IssueStatusKey:       IssueStatusData,
		IssuePriorityKey:     IssuePriorityData,
		IssueReproductionKey: IssueReproductionData,
	}

	// append all creators data into `pickData`
	allCreators := IssuePickerTemplateModel{}
	allCreators.Title = "Creators:"
	allCreators.DefaultValue = 0

	for _, eachUser := range allUsers {
		allCreators.Collection = append(allCreators.Collection, eachUser.Name)
	}
	c.issueTemplateData[IssueCreatorKey] = allCreators

	// append all assignor data into `pickData`
	allAssignors := IssuePickerTemplateModel{}
	allAssignors.Title = "Assignors:"
	allAssignors.DefaultValue = 0
	for _, eachAssignor := range allUsers {
		allAssignors.Collection = append(allAssignors.Collection, eachAssignor.Name)
	}
	c.issueTemplateData[IssueAssignorKey] = allAssignors
}

// initPageContent initial settings in current page
func (c *NewIssueController) initPageContent() {

	// generate page
	var htmlContent string
	var htmlContentSurfix string

	index := 0

	for key, value := range c.issueTemplateData {

		needRow := (index%3 == 0)

		if needRow {

			if len(htmlContentSurfix) > 0 {
				htmlContent += htmlContentSurfix
				htmlContentSurfix = ""
			}

			htmlContent += "<div class=\"form-group\">\n"
			htmlContent += "<div class=\"row\">\n"
		}

		htmlContent += "<div class=\"col-md-4\">\n"
		htmlContent += "<label class=\"right label-ff-standard\" style=\"width=100%\">" + key + "</label>\n"
		htmlContent += "<div class=\"btn-group\">\n"

		defaultValue := value.Collection[value.DefaultValue]
		htmlContent += "<button type=\"button\" class=\"btn btn-normal\" style=\"width=100%\">" + defaultValue + "</button>\n"
		htmlContent += "\n" +
			"<button type=\"button\" class=\"btn btn-normal dropdown-toggle\" data-toggle=\"dropdown\" aria-haspopup=\"true\" aria-expanded=\"false\">\n" +
			"<span class=\"caret\"></span>\n" +
			"<span class=\"sr-only\">Toggle Dropdown</span>\n" +
			"</button>\n"

		htmlContent += "<ul class=\"dropdown-menu\" style=\"height:15em;overflow-y:scroll;\">\n"
		for _, eachOption := range value.Collection {
			htmlContent += "<li><a href=\"#\">" + eachOption + "</a></li>\n"
		}
		htmlContent += "</ul>\n" // "<ul class=\"dropdown-menu\">\n"

		htmlContent += "</div>\n" // "<div class=\"btn-group\">"
		htmlContent += "</div>\n" // "<div class=\"col-md-4\">"

		if needRow {

			if len(htmlContentSurfix) > 0 {
				htmlContent += htmlContentSurfix
				htmlContentSurfix = ""
			}

			htmlContentSurfix += "</div>\n"
			htmlContentSurfix += "</div>\n"
		}

		index++
	}

	if len(htmlContentSurfix) > 0 {
		htmlContent += htmlContentSurfix
		htmlContentSurfix = ""
	}

	c.Data[constants.KeyIssueInitValue] = htmlContent
}
