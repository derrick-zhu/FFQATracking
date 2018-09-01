package controllers

import (
	"FFQATracking/constants"
)

// NewIssueController base issue create page
type NewIssueController struct {
	FFBaseController
}

// Get for handle new issue controller GET request
func (c *NewIssueController) Get() {
	c.FFBaseController.Get()
	c.Data[constants.KeyIsIssueList] = 1

	c.TplName = "newIssue.html"
}
