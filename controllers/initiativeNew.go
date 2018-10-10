package controllers

import (
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

const (
	modelPropertySectionKey string = "modelPropertySection"
)

// InitiativeNewController controller for create initiative
type InitiativeNewController struct {
	FFBaseController
}

func (c *InitiativeNewController) Get() {

	c.FFBaseController.Get()

	c.initCommonVar()
	beego.Error(c.Data)
	c.TplName = "initiativeNew.html"
}

/**
private helpers
*/

func (c *InitiativeNewController) initCommonVar() {

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
	)

	initiativeProperties = append(
		initiativeProperties,
		models.DataTextareaTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Description:",
				Identifier: "description",
				Type:       models.TextArea,
			},
			DefaultValue: "",
			Value:        "",
		},
	)

	initiativeProperties = append(
		initiativeProperties,
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
	)

	initiativeProperties = append(
		initiativeProperties,
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
	)

	c.Data[modelPropertySectionKey] = initiativeProperties
}
