package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

const (
	// BugsTable db table name for bugmodel
	BugsTable string = "bugmodel"
)

// BugStatus bug status type
type BugStatus string

const (
	BugNew        BugStatus = "New"
	BugFixed      BugStatus = "Fixed"
	BugReopen     BugStatus = "Reopen"
	BugConfirm    BugStatus = "Confirm"
	BugClose      BugStatus = "Close"
	BugNotABug    BugStatus = "Not a bug"
	BugWillNotFix BugStatus = "Will not fix"
	BugDelay      BugStatus = "Delay"
	BugMustBeFix  BugStatus = "Must be fix"
)

// PriorityStatus bug's priority status
type PriorityStatus int64

const (
	PriorityUrgent     PriorityStatus = 0
	PriorityImportant  PriorityStatus = 1
	PriorityHigh       PriorityStatus = 2
	PriorityMiddle     PriorityStatus = 3
	PriorityLow        PriorityStatus = 4
	PriorityQuestion   PriorityStatus = 5
	PrioritySuggestion PriorityStatus = 6
)

// BugModel the model of bug
type BugModel struct {
	ID              IndexType      `orm:"index"`      // index
	Title           string         `orm:"size(512)"`  // bug title
	Description     string         `orm:"size(4096)"` // description about bug
	Version         string         `orm:"index"`      // test version number
	Source          string         // source feature request
	Target          string         `orm:"index"` // target milestone
	DevPeriod       string         `orm:"index"` // sprint
	SolveDate       time.Time      // date solving
	CreateDate      time.Time      // date creating
	Status          BugStatus      `orm:"index"` // bug current status
	Priority        PriorityStatus `orm:"index"` // bug's priority type
	Creator         IndexType      `orm:"index"` // bug's founder
	Assignor        IndexType      `orm:"index"` // who should solve this bug
	Reproducibility int            // 重现概率 0~100
}

func init() {
	orm.RegisterModel(new(BugModel))
}

// TableName for beego using
func (c *BugModel) TableName() string {
	return BugsTable
}

// AddBug insert new bug
func AddBug(title, description string, creatorID IndexType) (*BugModel, error) {
	pBug := &BugModel{
		Title:       title,
		Description: description,
		Creator:     creatorID,
		CreateDate:  time.Now(),
		Status:      BugNew,
	}

	o, _ := GetQuerySeterWithTable(BugsTable)
	_, err := o.Insert(pBug)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return pBug, nil
}

// BugWithID fetch bug with id
func BugWithID(id IndexType) (*BugModel, error) {

	pbug := &BugModel{ID: id}

	_, qs := GetQuerySeterWithTable(BugsTable)
	filterErr := qs.Filter("id", id).One(pbug)
	if filterErr != nil {

		beego.Error(filterErr)

		return nil, filterErr
	}

	return pbug, nil
}

// BugsWithRange fetch bug data with range [lower, lower + count)
func BugsWithRange(lower, count int) ([]*BugModel, error) {

	var result []*BugModel
	var err error
	var rawResult orm.RawSeter

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", BugsTable, count, lower)
	rawResult = o.Raw(sqlQuery)

	_, err = rawResult.QueryRows(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AllBugs fetch all bugs
func AllBugs() ([]*BugModel, error) {

	return BugsWithRange(0, -1)
}

// UpdateBug update bug model data
func UpdateBug(id IndexType, params map[string]interface{}) error {

	_, qs := GetQuerySeterWithTable(BugsTable)
	_, err := BugWithID(id)

	if err != nil {
		beego.Error(err)

		return err
	}

	beego.Info(params)

	_, err = qs.Filter("id", id).Update(params)
	if err != nil {
		beego.Error(err)

		return err
	}

	return nil
}

// DeleteBug delete bug with id
func DeleteBug(id IndexType) error {

	o, _ := GetQuerySeterWithTable(BugsTable)

	bug := &BugModel{ID: id}
	_, err := o.Delete(bug)

	return err
}
