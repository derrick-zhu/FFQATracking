package controllers

import (
	"FFQATracking/models"
	"FFQATracking/utils"

	"github.com/astaxie/beego"
)

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
