package controllers

import (
	"FFQATracking/models"
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

}
