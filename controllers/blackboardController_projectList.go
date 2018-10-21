package controllers

import (
	"FFQATracking/models"
)

const (
	// AllInitiativesListConst ...
	AllInitiativesListConst = "allInitiativesList"
)

// InitiativeTableItemModel ...
type InitiativeTableItemModel struct {
	models.InitiativeModel

	CreatorName  string
	AssignorName string
}

func (c *InitiativeTableItemModel) initWith(initiative *models.InitiativeModel) {
	c.InitiativeModel = models.InitiativeModel{
		ID:          initiative.ID,
		Name:        initiative.Name,
		Description: initiative.Description,
		Creator:     initiative.Creator,
		Assignor:    initiative.Assignor,
		StartDate:   initiative.StartDate,
		EndDate:     initiative.EndDate,
	}
}

func (c *InitiativeTableItemModel) setupInitiativeWithUserList(allUsers *[]models.AccountModel) {

	for _, eachUser := range *allUsers {

		if c.Creator == eachUser.ID {
			c.CreatorName = eachUser.Name
		}

		if c.Assignor == eachUser.ID {
			c.AssignorName = eachUser.Name
		}
	}
}

func (c *BlackboardController) initProjectListVar(allUsers *[]models.AccountModel, allProjects *[]models.InitiativeModel) {

	var allProjectItems = []InitiativeTableItemModel{}
	for _, eachInit := range *allProjects {

		initiativeItemData := InitiativeTableItemModel{}
		initiativeItemData.initWith(&eachInit)
		initiativeItemData.setupInitiativeWithUserList(allUsers)

		allProjectItems = append(allProjectItems, initiativeItemData)
	}
	c.Data[AllInitiativesListConst] = allProjectItems
}
