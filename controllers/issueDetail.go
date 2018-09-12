package controllers

import (
	"FFQATracking/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type IssueDetailController struct {
	FFBaseController

	issueDetailData TIssueNewCollectionType
	issueID         int64
	currentIssue    models.BugModel
}

func (c *IssueDetailController) Get() {
	c.FFBaseController.Get()

	beego.Info(c.Ctx.Input)

	var err error
	c.issueID, err = strconv.ParseInt(c.Ctx.Input.Param(":issue"), 10, 64)
	if err != nil {
		beego.Error(err)
		c.Redirect("/issuelist", 302)
		return
	}
	beego.Info(fmt.Sprintf("issue Id: %d", c.issueID))

	c.initVariables()

	c.TplName = "issueDetail.html"
}

func (c *IssueDetailController) initVariables() {

	c.currentIssue = models.BugWithID(c.issueID)

	statusData := IssueStatusData
	statusData.DefaultValue = c.currentIssue.Status

	// c.issueDetailData = TIssueNewCollectionType{
	// 	IssueStatusData,
	// 	IssuePriorityData,
	// 	IssueReproductionData,
	// }

	// allUsers, err := models.AllAccounts()
	// if err != nil {
	// 	beego.Error(err)
	// }

	// allCreators := IssuePickerTemplateModel{}
	// allCreators.Title = IssueCreatorKey
	// allCreators.Identifier = fmt.Sprintf("%s%s", issueIDPrefix, allCreators.Title)
	// allCreators.DefaultValue =
	// allCreators.Collection = allUsers
}
