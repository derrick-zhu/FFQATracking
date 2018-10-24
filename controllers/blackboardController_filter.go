package controllers

import (
	"FFQATracking/models"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	allFilterConst = "allFilters"
)

// FilterChanged handle picker's value changed event
func (c *BlackboardController) FilterChanged() {
	c.FFBaseController.Get()

	selInitiativeID, _ := strconv.ParseInt(c.GetString("initiative_id"), 10, 64)
	selMilestoneID, _ := strconv.ParseInt(c.GetString("milestone_id"), 10, 64)

	c.commonInitForGet(selInitiativeID, selMilestoneID)

	c.ServeJSON()
}

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

	if len(allInitiativeVar) == 0 {
		_foo := models.ZeroInitiative()
		_foo.Name = "-- No filter --"
		allInitiativeVar = append(allInitiativeVar, _foo)
	}

	var selIdx int
	for idx, eachInit := range *allInitiatives {
		allInitiativeVar = append(allInitiativeVar, eachInit)
		if eachInit.ID == selID {
			selIdx = idx
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
	// default item - "No data"
	if len(msFilterResult) == 0 {
		_foo := models.ZeroMilestone()
		_foo.Name = "-- No data --"
		msFilterResult = append(msFilterResult, _foo)
	}

	var selIdx int
	for idx, eachMS := range *msFilterArrs {
		msFilterResult = append(msFilterResult, eachMS)
		if eachMS.ID == selMSID {
			selIdx = idx
		}
	}

	return &msFilterResult, selIdx
}
