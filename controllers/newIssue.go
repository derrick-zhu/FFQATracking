package controllers

import (
	"FFQATracking/constants"
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

// PickerTemplateModel class template
type PickerTemplateModel struct {
	Title        string
	DefaultValue int
	Collection   []string
}

type TIssueNewCollectionType map[string]PickerTemplateModel

var IssueStatusData PickerTemplateModel = PickerTemplateModel{
	Title:        IssueStatusKey,
	DefaultValue: 0,
	Collection:   []string{"New", "Fixed", "Reopen", "Confirm", "Close", "Not a bug", "Will not fix", "Delay", "Must be fix"},
}

var IssuePriorityData PickerTemplateModel = PickerTemplateModel{
	Title:        IssuePriorityKey,
	DefaultValue: 0,
	Collection:   []string{"Urgent", "Important", "High", "Middle", "Low", "Question", "Suggestion"},
}

var IssueReproductionData PickerTemplateModel = PickerTemplateModel{
	Title:        IssueReproductionKey,
	DefaultValue: 0,
	Collection:   []string{"100%", "80%", "60%", "40%", "20%"},
}

// NewIssueController base issue create page
type NewIssueController struct {
	FFBaseController
}

// Get for handle new issue controller GET request
func (c *NewIssueController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.Title] = "New Issue"
	c.Data[constants.KeyIsIssueList] = 1

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

// initPageContent initial settings in current page
func (c *NewIssueController) initPageContent() {

	var pickerData = TIssueNewCollectionType{

		IssueStatusKey:       IssueStatusData,
		IssuePriorityKey:     IssuePriorityData,
		IssueReproductionKey: IssueReproductionData,
	}

	var htmlContentPrefix string
	var htmlContentSurfix string

	index := 0

	for key, value := range pickerData {

		needRow := (index%3 == 0)

		if needRow {

			if len(htmlContentSurfix) > 0 {
				htmlContentPrefix += htmlContentSurfix
				htmlContentSurfix = ""
			}

			htmlContentPrefix += "<div class=\"form-group\">\n"
			htmlContentPrefix += "<div class=\"row\">\n"
		}

		htmlContentPrefix += "<div class=\"col-md-4\">\n"
		htmlContentPrefix += "<label class=\"right label-ff-standard\">" + key + "</label>\n"
		htmlContentPrefix += "<div class=\"btn-group\">\n"

		defaultValue := value.Collection[value.DefaultValue]
		htmlContentPrefix += "<button type=\"button\" class=\"btn btn-normal\">" + defaultValue + "</button>\n"
		htmlContentPrefix += "\n" +
			"<button type=\"button\" class=\"btn btn-normal dropdown-toggle\" data-toggle=\"dropdown\" aria-haspopup=\"true\" aria-expanded=\"false\">\n" +
			"<span class=\"caret\"></span>\n" +
			"<span class=\"sr-only\">Toggle Dropdown</span>\n" +
			"</button>\n"

		htmlContentPrefix += "<ul class=\"dropdown-menu\">\n"
		for _, eachOption := range value.Collection {
			htmlContentPrefix += "<li><a href=\"#\">" + eachOption + "</a></li>\n"
		}
		htmlContentPrefix += "</ul>\n" // "<ul class=\"dropdown-menu\">\n"

		htmlContentPrefix += "</div>\n" // "<div class=\"btn-group\">"
		htmlContentPrefix += "</div>\n" // "<div class=\"col-md-4\">"

		if needRow {

			if len(htmlContentSurfix) > 0 {
				htmlContentPrefix += htmlContentSurfix
				htmlContentSurfix = ""
			}

			htmlContentSurfix += "</div>\n"
			htmlContentSurfix += "</div>\n"
		}

		index++
	}

	if len(htmlContentSurfix) > 0 {
		htmlContentPrefix += htmlContentSurfix
		htmlContentSurfix = ""
	}

	c.Data[constants.KeyIssueInitValue] = pickerData
	c.Data["test"] = htmlContentPrefix
}
