package controllers

import (
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

const (
	allFilterConst = "allFilters"
)

func (c *BlackboardController) initFilterVars(allUser *[]models.AccountModel, allInitiatives *[]models.InitiativeModel) {

	const defIndex int64 = 0
	var initID int64

	// all initiatives
	var initiativeFilterArrs = []interface{}{}
	var allInitiativeVar = []models.VarModelProtocol{}

	if len(allInitiativeVar) == 0 {
		_foo := models.ZeroInitiative()
		_foo.Name = "-- No filter --"
		allInitiativeVar = append(allInitiativeVar, _foo)
	}

	for idx, eachInit := range *allInitiatives {
		if int64(idx) == defIndex {
			initID = eachInit.ID
		}
		allInitiativeVar = append(allInitiativeVar, eachInit)
	}

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

	initiativeFilterArrs = append(
		initiativeFilterArrs,
		// initiatives
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Initiatives",
				Identifier: "initiatives",
				Type:       models.Number,
			},
			DefaultValue: defIndex,
			Value:        0,
			Collection:   allInitiativeVar,
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

	c.Data[allFilterConst] = initiativeFilterArrs
}

func (c *BlackboardController) fetchMilestoneFilterWithInitiativeID(initID, offset, count int64) (*[]models.MilestoneModel, error) {

	result, err := models.MilestonesWithInitiative(initID, offset, count)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return result, nil
}
