package controllers

import (
	"FFQATracking/models"

	"github.com/astaxie/beego"
)

const (
	allFilterConst         = "allFilters"
	currentInitiativeConst = "currentInitiative"
	currentMilestoneConst  = "currentMilestone"
)

func (c *BlackboardController) initFilterVars(selectedProjID, selectedMilestoneID int64, allUser *[]models.AccountModel, allInitiatives *[]models.InitiativeModel) {

	// all filters
	var blackboardFilterArrs = []interface{}{}

	// all initiative filters
	allInitiativeVar, defInitIndex := c.generateInitiativeFilters(allInitiatives, selectedProjID)
	// all milestone filters
	allMilestoneFilterVar, defMSIndex := c.generateMilestoneFilters(selectedProjID, selectedMilestoneID)

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
			ValueChanged: models.JSCommandModel{
				ID: "initiativePickerChanged",
			},
			DefaultValue: int64(defInitIndex),
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
					ID:   "#bbNewMilestoneModal",
					Name: "New Issue",
				},
			},
			ValueChanged: models.JSCommandModel{
				ID: "milestonePickerValueChanged",
			},
			DefaultValue: int64(defMSIndex),
			Value:        0,
			Collection:   *allMilestoneFilterVar,
		},
	)

	c.Data[allFilterConst] = blackboardFilterArrs
}

// private helpers

func (c *BlackboardController) indexOfInitiativeWith(projID int64, allInitiatives *[]models.InitiativeModel) int {
	for idx, init := range *allInitiatives {
		if init.ID == projID {
			return idx
		}
	}
	return -1
}

func (c *BlackboardController) indexOfMilestoneWith(msID int64, allMilestones *[]models.MilestoneModel) int {
	for idx, ms := range *allMilestones {
		if ms.ID == msID {
			return idx
		}
	}
	return -1
}

func (c *BlackboardController) generateInitiativeFilters(allInitiatives *[]models.InitiativeModel, selID int64) (*[]models.VarModelProtocol, int) {

	var allInitiativeVar = []models.VarModelProtocol{}

	// fill with all options, including EMPTY one.
	if len(allInitiativeVar) == 0 {
		_foo := models.ZeroInitiative()
		_foo.Name = "-- All --"
		allInitiativeVar = append(allInitiativeVar, _foo)
	}

	for _, eachInit := range *allInitiatives {
		allInitiativeVar = append(allInitiativeVar, eachInit)
	}

	// then, find the selected item's index
	var selIdx int
	for idx, eachInit := range allInitiativeVar {
		if eachInit.(models.InitiativeModel).ID == selID {
			selIdx = idx
			break
		}
	}

	return &allInitiativeVar, selIdx
}

func (c *BlackboardController) generateMilestoneFilters(selInitID int64, selMSID int64) (*[]models.VarModelProtocol, int) {

	msFilterArrs, err := c.fetchMilestoneFilterWithInitiativeID(selInitID, 0, -1)
	if err != nil {
		beego.Error(err)
		err = nil
		msFilterArrs = &[]models.MilestoneModel{}
	}

	msFilterResult := []models.VarModelProtocol{}
	// fill with all options, including EMPTY one.
	// default item - "No data"
	if len(msFilterResult) == 0 {
		_foo := models.ZeroMilestone()
		_foo.Name = "-- All --"
		msFilterResult = append(msFilterResult, _foo)
	}

	for _, eachMS := range *msFilterArrs {
		msFilterResult = append(msFilterResult, eachMS)
	}

	// then, find the selected item's index
	var selIdx int
	for idx, eachMS := range msFilterResult {
		if eachMS.(models.MilestoneModel).ID == selMSID {
			selIdx = idx
		}
	}

	return &msFilterResult, selIdx
}
