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
	// all filters
	var blackboardFilterArrs = []interface{}{}

	// all initiative filters
	allInitiativeVar := c.generateInitiativeFilters(defIndex, &initID, allInitiatives)
	// all milestone filters
	allMilestoneFilterVar := c.generateMilestoneFilters(initID)

	blackboardFilterArrs = append(
		blackboardFilterArrs,
		// initiatives
		models.DataPickerTemplateModel{
			DataBaseTemplateModel: models.DataBaseTemplateModel{
				Title:      "Initiatives",
				Identifier: "initiatives",
				Type:       models.Number,
				JSCmd: models.JSCommandModel{
					ID:   "#bbNewInitiativeModal",
					Name: "New Project",
				},
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
				JSCmd: models.JSCommandModel{
					ID:   "#bbNewIssueModal",
					Name: "New Issue",
				},
			},
			DefaultValue: 0,
			Value:        0,
			Collection:   *allMilestoneFilterVar,
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

func (c *BlackboardController) generateMilestoneFilters(defInitID int64) *[]models.VarModelProtocol {

	msFilterArrs, err := c.fetchMilestoneFilterWithInitiativeID(defInitID, 0, -1)
	if err != nil {
		beego.Error(err)
		err = nil
		msFilterArrs = &[]models.MilestoneModel{}
	}

	msFilterResult := []models.VarModelProtocol{}
	// default item - "No data"
	if len(msFilterResult) == 0 {
		_foo := models.ZeroMilestone()
		_foo.Name = "-- No data --"
		msFilterResult = append(msFilterResult, _foo)
	}

	for _, eachMS := range *msFilterArrs {
		msFilterResult = append(msFilterResult, eachMS)
	}

	return &msFilterResult
}
