package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	initiativeModelTblName string = "initiative"
)

// InitiativeModel db struct def
type InitiativeModel struct {
	ID          int64  `orm:"pk;auto;index"`
	Name        string `orm:"size(256)"`
	Description string `orm:"size(4096)"`
	Creator     int64
	Assignor    int64
}

func init() {
	orm.RegisterModel(new(InitiativeModel))
}

// TableName db table name
func (c *InitiativeModel) TableName() string {
	return initiativeModelTblName
}

// NewInitiative new and insert initiative data
func NewInitiative(name, desc string, creatorID int64) (*InitiativeModel, error) {

	aNewInit := &InitiativeModel{
		Name:        name,
		Description: desc,
	}

	o := GetOrmObject()
	if _, err := o.Insert(aNewInit); err != nil {
		beego.Error(err)
		o.Rollback()
		return nil, err
	}

	return aNewInit, nil
}

// RemoveInitiative delete a initiative obj with its id
func RemoveInitiative(initiativeID int64) error {

	o, _ := GetQuerySeterWithTable(initiativeModelTblName)

	var err error
	aInitiative := &InitiativeModel{ID: initiativeID}
	if _, err = o.Delete(aInitiative); err != nil {
		beego.Error(err)
		o.Rollback()
	}

	return err
}

// InitiativeUpdate update initiative with its new properties
func InitiativeUpdate(newInitiative *InitiativeModel) error {

	o, _ := GetQuerySeterWithTable(initiativeModelTblName)

	if _, err := o.Update(newInitiative); err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}

	return nil
}

// Initiatives fetch initiative data
func Initiatives(low, count int64) (*[]InitiativeModel, error) {

	var results = &[]InitiativeModel{}
	var err error
	var rawSeter orm.RawSeter

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FORM %s LIMIT %d OFFSET %d", initiativeModelTblName, count, low)
	rawSeter = o.Raw(sqlQuery)

	if _, err = rawSeter.QueryRows(results); err != nil {
		beego.Error(err)
		return nil, err
	}

	return results, nil
}

// InitiativeWithID fetch initiative data with its id
func InitiativeWithID(initiativeID int64) (*InitiativeModel, error) {

	pInitiative := &InitiativeModel{ID: initiativeID}
	o, _ := GetQuerySeterWithTable(initiativeModelTblName)

	if err := o.Read(pInitiative); err != nil {
		beego.Error(err)
		return nil, err
	}
	return pInitiative, nil
}

// InitiativesWithCreator fetch initiative which created by someone
func InitiativesWithCreator(creatorID, low, count int64) (*[]InitiativeModel, error) {

	var result = &[]InitiativeModel{}
	var rawSeter orm.RawSeter

	o := GetOrmObject()
	sqlQuery :=
		fmt.Sprintf("SELECT * FROM %s WHERE creator == %d LIMIT %d offset %d",
			initiativeModelTblName,
			creatorID,
			count,
			low)
	rawSeter = o.Raw(sqlQuery)

	if _, err := rawSeter.QueryRows(result); err != nil {
		beego.Error(err)
		return nil, err
	}

	return result, nil
}

// InitiativesWithAssignor fetch initiative which assigned to somebody
func InitiativesWithAssignor(assignorID, low, count int64) (*[]InitiativeModel, error) {

	var result = &[]InitiativeModel{}
	var rawSeter orm.RawSeter

	o := GetOrmObject()
	sqlQuery :=
		fmt.Sprintf("SELECT * FROM %s WHERE assignor == %d LIMIT %d offset %d",
			initiativeModelTblName,
			assignorID,
			count,
			low)
	rawSeter = o.Raw(sqlQuery)

	if _, err := rawSeter.QueryRows(result); err != nil {
		beego.Error(err)
		return nil, err
	}

	return result, nil
}
