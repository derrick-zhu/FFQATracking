package controllers

import (
	"FFQATracking/models"
)

const (
	allFilterConst = "allFilters"
)

func (c *BlackboardController) initFilterVars(allUser *[]models.AccountModel, allInitiatives *[]models.InitiativeModel) {

	var initiativeFilterArrs = []interface{}{}
	var allInitiativeVar = []models.VarModelProtocol{}
	for _, eachInit := range *allInitiatives {
		allInitiativeVar = append(allInitiativeVar, eachInit)
	}

	initiativeFilterArrs = append(
		initiativeFilterArrs,
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Initiatives",
				Identifier: "initiatives",
				Type:       models.Number,
			},
			DefaultValue: 0,
			Value:        0,
			Collection:   allInitiativeVar,
		},
	)
	c.Data[allFilterConst] = initiativeFilterArrs
}
