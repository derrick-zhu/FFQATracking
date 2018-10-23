package controllers

import (
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

const (
	allFilterConst = "allFilters"
)

func (c *BlackboardController) initFilterVars(allUser *[]models.AccountModel, allInitiatives *[]models.InitiativeModel) {

	const defIndex int = 0
	var initID int64

	// all initiatives
	var blackboardFilterArrs = []interface{}{}

	allInitiativeVar := c.generateInitiativeFilters(defIndex, &initID, allInitiatives)

	// all milestone filters
	msFilterArrs, err := c.fetchMilestoneFilterWithInitiativeID(initID, 0, -1)
	if err != nil {
		beego.Error(err)
		err = nil
		msFilterArrs = &[]models.MilestoneModel{}
	}

	msFilterVar := []models.VarModelProtocol{}

	if len(msFilterVar) == 0 {
		_foo := models.ZeroMilestone()
		_foo.Name = "-- No data --"
		msFilterVar = append(msFilterVar, _foo)
	}

	for _, eachMS := range *msFilterArrs {
		msFilterVar = append(msFilterVar, eachMS)
	}

	blackboardFilterArrs = append(
		blackboardFilterArrs,
		// initiatives
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Initiatives",
				Identifier: "initiatives",
				Type:       models.Number,
			},
			DefaultValue: int64(defIndex),
			Value:        0,
			Collection:   *allInitiativeVar,
		},
		// milestone
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Version",
				Identifier: "versions",
				Type:       models.Number,
			},
			DefaultValue: 0,
			Value:        0,
			Collection:   msFilterVar,
		},
	)

	c.Data[allFilterConst] = blackboardFilterArrs
}

// private helpers

func (c *BlackboardController) generateInitiativeFilters(defIndex int, initID *int64, allInitiatives *[]models.InitiativeModel) *[]models.VarModelProtocol {

	var allInitiativeVar = []models.VarModelProtocol{}

	if len(allInitiativeVar) == 0 {
		_foo := models.ZeroInitiative()
		_foo.Name = "-- No filter --"
		allInitiativeVar = append(allInitiativeVar, _foo)
	}

	for idx, eachInit := range *allInitiatives {
		if idx == defIndex {
			*initID = eachInit.ID
		}
		allInitiativeVar = append(allInitiativeVar, eachInit)
	}

	return &allInitiativeVar
}
