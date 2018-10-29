package controllers

import (
	"FFQATracking/biz"
	"FFQATracking/models"
	"FFQATracking/utils"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	newVersionFieldConst = "newVersionField"
)

func (c *BlackboardController) initNewMilestoneVar() {

	var fieldNewMilestone = models.DataFieldTemplateModel{
		DataBaseTemplateModel: models.DataBaseTemplateModel{
			Title:      "Version",
			Identifier: "version",
			Type:       models.TextField,
		},
		DefaultValue: "",
		Value:        "",
	}

	c.Data[newVersionFieldConst] = fieldNewMilestone
}

// SubmitNewMilestone ...
func (c *BlackboardController) SubmitNewMilestone() {

	c.FFBaseController.Post()

	var err error
	var selProject int64
	var newMilestone string
	var acc *models.AccountModel
	var ms *models.MilestoneModel

	for {
		if biz.HadLogin(c.Ctx) == false {
			utils.MakeRedirectURL(&c.Data, 302, "/login", "")
			break
		}

		if acc, err = biz.CurrentAccount(c.Ctx); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/blackboard", "")
			break
		}

		if selProject, err = strconv.ParseInt(c.GetString("project", "-1"), 10, 64); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		if newMilestone = c.GetString("milestone", ""); len(newMilestone) <= 0 {
			beego.Error("invalid milestone value.")
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		if ms, err = models.AddMilestone(newMilestone, selProject, acc.ID); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		paramForJs := models.GOCommandModel{
			Param: map[string]interface{}{
				"proj":    selProject,
				"msid":    ms.ID,
				"creator": acc.ID,
			},
		}

		utils.MakeRedirectURLWithUserInfo(&c.Data, 302, "#", "", paramForJs)
		break
	}

	c.ServeJSON()
}
