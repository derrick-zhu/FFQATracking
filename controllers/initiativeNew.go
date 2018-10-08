package controllers

import (
	"FFQATracking/models"
)

const (
	modelPropertySectionKey string = "modelPropertySection"
)

/**
page's basic data
*/

var initiativeProperties = []interface{}{
	models.DataFieldTemplateModel{
		BaseDataTemplateModel: models.BaseDataTemplateModel{
			Title:      "Title:",
			Identifier: "title",
			Type:       models.TextField,
		},
		DefaultValue: "",
		Value:        "",
	},

	models.DataFieldTemplateModel{
		BaseDataTemplateModel: models.BaseDataTemplateModel{
			Title:      "Description:",
			Identifier: "description",
			Type:       models.TextArea,
		},
		DefaultValue: "",
		Value:        "",
	},

	models.DataPickerTemplateModel{
		BaseDataTemplateModel: models.BaseDataTemplateModel{
			Title:      "Creator:",
			Identifier: "creator",
			Type:       models.Number,
		},
		DefaultValue: 0,
		Value:        0,
	},

	models.DataPickerTemplateModel{
		BaseDataTemplateModel: models.BaseDataTemplateModel{
			Title:      "Assignor:",
			Identifier: "assignor",
			Type:       models.Number,
		},
		DefaultValue: 0,
		Value:        0,
	},
}

// InitiativeNewController controller for create initiative
type InitiativeNewController struct {
	FFBaseController
}

func (c *InitiativeNewController) Get() {

	c.FFBaseController.Get()

	c.initCommonVar()
	c.TplName = "initiativeNew.html"
}

/**
private helpers
*/

func (c *InitiativeNewController) initCommonVar() {

	// var allUsers *[]models.AccountModel
	// var err error

	// for true {

	// 	if allUsers, err = models.AllAccounts(); err != nil {
	// 		beego.Error(err)
	// 		break
	// 	}

	// }

	c.Data[modelPropertySectionKey] = initiativeProperties
}
