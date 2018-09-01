package controllers

import (
	"FFQATracking/constants"
)

// IssueListController the class of issue list
type IssueListController struct {
	FFBaseController
}

// Get for handling issue list page HTTP GET request
func (c *IssueListController) Get() {
	c.FFBaseController.Get()

	c.Data[constants.KeyIsIssueList] = 1

	c.TplName = "issueList.html"
}
