package models

import (
	"FFQATracking/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	milestoneTableNameConst = "milestones"
)

// MilestoneModel ...
type MilestoneModel struct {
	ID           int64  `orm:"auto;index;pk"`
	Name         string `orm:"size(512)"`
	Creator      int64
	Date         int64
	InitiativeID int64 `orm:"index"`
}

// Type implement VarModelProtocol
func (c MilestoneModel) Type() int64 { return c.ID }

// Desc implement VarModelProtocol
func (c MilestoneModel) Desc() string { return c.Name }

//ZeroMilestone implement EmptyDataProtocol, create an empty milestone model data,
func ZeroMilestone() *MilestoneModel {
	return &MilestoneModel{ID: -1}
}

func init() {
	orm.RegisterModel(new(MilestoneModel))
}

// TableName implement db table name protocol
func (c *MilestoneModel) TableName() string {
	return milestoneTableNameConst
}

// AddMilestone ...
func AddMilestone(name string, initiativeID, creatorID int64) (*MilestoneModel, error) {
	result := &MilestoneModel{
		Name:         name,
		InitiativeID: initiativeID,
		Creator:      creatorID,
		Date:         utils.TimeTickSince1970(),
	}

	o, _ := GetQuerySeterWithTable(milestoneTableNameConst)
	if _, err := o.Insert(result); err != nil {
		beego.Error(err)
		o.Rollback()

		return nil, err
	}

	return result, nil
}

// RemoveMilestone remove milestone with its id
func RemoveMilestone(id int64) error {

	ms := &MilestoneModel{ID: id}
	o, _ := GetQuerySeterWithTable(milestoneTableNameConst)
	if _, err := o.Delete(ms); err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}
	return nil
}

// UpdateMilestone update milestone data
func UpdateMilestone(newMS *MilestoneModel) error {

	o, _ := GetQuerySeterWithTable(milestoneTableNameConst)
	if _, err := o.Update(newMS); err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}
	return nil
}

// MilestoneWithID fetch milestone data from db
func MilestoneWithID(id int64) (*MilestoneModel, error) {
	result := &MilestoneModel{ID: id}
	o, _ := GetQuerySeterWithTable(milestoneTableNameConst)
	if err := o.Read(result); err != nil {
		beego.Error(err)
		return nil, err
	}
	return result, nil
}

// MilestonesWithInitiative ...
func MilestonesWithInitiative(initiativeID, offset, count int64) (*[]MilestoneModel, error) {

	result := &[]MilestoneModel{}
	var err error

	o := GetOrmObject()
	sqlquery := fmt.Sprintf("SELECT * FROM %s WHERE initiative_i_d = %d LIMIT %d OFFSET %d", milestoneTableNameConst, initiativeID, count, offset)
	rawResult := o.Raw(sqlquery)

	if _, err = rawResult.QueryRows(result); err != nil {
		beego.Error(err)
		return nil, err
	}
	return result, nil
}
