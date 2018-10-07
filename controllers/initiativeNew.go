package controllers

import (
	"FFQATracking/models"
	"FFQATracking/models/private"
)

/**
page's basic data
*/

var initiativeProperties = []interface{}{
	models.DataFieldTemplateModel{
		BaseDataTemplateModel: private.BaseDataTemplateModel{
			Title:      "Title:",
			Identifier: "",
			Type:       private.TextField,
		},
		DefaultValue: "",
		Value:        "",
	},

	models.DataFieldTemplateModel{
		BaseDataTemplateModel: private.BaseDataTemplateModel{
			Title:      "Description:",
			Identifier: "",
			Type:       private.TextArea,
		},
		DefaultValue: "",
		Value:        "",
	},

	models.DataPickerTemplateModel{
		BaseDataTemplateModel: private.BaseDataTemplateModel{
			Title:      "Creator:",
			Identifier: "",
			Type:       private.Number,
		},
		DefaultValue: 0,
		Value:        0,
	},

	models.DataPickerTemplateModel{
		BaseDataTemplateModel: private.BaseDataTemplateModel{
			Title:      "Assignor:",
			Identifier: "",
			Type:       private.Number,
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

	c.TplName = "initiativeNew.html"
}

/**
private helpers
*/

func (c *InitiativeNewController) initCommonVar() {

}
