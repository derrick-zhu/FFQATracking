package controllers

import (
	"FFQATracking/models"
	"FFQATracking/utils"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	modelPropertySectionKey = "modelPropertySection"
)

// SubmitNewProject handle POST http request for creating project
func (c *BlackboardController) SubmitNewProject() {

	c.FFBaseController.Post()

	beego.Info(c.Input())
	beego.Info(c.GetString("title"))

	var err error
	var creator, assignor, startDate, endDate int64

	for {
		title := c.GetString("title")
		description := c.GetString("description")

		if creator, err = strconv.ParseInt(c.GetString("creator"), 10, 64); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		if assignor, err = strconv.ParseInt(c.GetString("assignor"), 10, 64); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		if startDate, err = strconv.ParseInt(c.GetString("startDate"), 10, 64); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		if endDate, err = strconv.ParseInt(c.GetString("endDate"), 10, 64); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		if _, err = models.NewInitiative(title, description, creator, assignor, startDate, endDate); err != nil {
			beego.Error(err)
			utils.MakeRedirectURL(&c.Data, 500, "/#", "")
			break
		}

		utils.MakeRedirectURL(&c.Data, 302, "/blackboard", "")
		break
	}

	c.ServeJSON()
}

func (c *BlackboardController) initNewInitiativeVar() {
	var initiativeProperties = []interface{}{}
	var allUsers *[]models.AccountModel
	var allUsersVar = []models.VarModelProtocol{}
	var err error

	if allUsers, err = models.AllAccounts(); err != nil {
		beego.Error(err)
	}

	for _, v := range *allUsers {
		allUsersVar = append(allUsersVar, v)
	}

	initiativeProperties = append(
		initiativeProperties,

		models.DataFieldTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Title:",
				Identifier: "title",
				Type:       models.TextField,
			},
			DefaultValue: "",
			Value:        "",
		},

		models.DataTextareaTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Description:",
				Identifier: "description",
				Type:       models.TextArea,
			},
			DefaultValue: "",
			Value:        "",
		},

		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Creator:",
				Identifier: "creator",
				Type:       models.Number,
			},
			DefaultValue: 0,
			Value:        0,
			Collection:   allUsersVar,
		},

		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Assignor:",
				Identifier: "assignor",
				Type:       models.Number,
			},
			DefaultValue: 0,
			Value:        0,
			Collection:   allUsersVar,
		},
		// start date
		models.DataDatePickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Start Date:",
				Identifier: "startDate",
				Type:       models.Date,
			},
			DefaultValue: utils.TimeTickSince1970(),
			Value:        utils.TimeTickSince1970(),
		},
		// end date
		models.DataDatePickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "End Date:",
				Identifier: "endDate",
				Type:       models.Date,
			},
			DefaultValue: utils.TimeTickSince1970(),
			Value:        utils.TimeTickSince1970(),
		},
	)

	c.Data[modelPropertySectionKey] = initiativeProperties
}
