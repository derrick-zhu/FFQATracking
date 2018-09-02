package models

import (
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
	psUrgent     PriorityStatus = 0
	psImportant  PriorityStatus = 1
	psHigh       PriorityStatus = 2
	psMiddle     PriorityStatus = 3
	psLow        PriorityStatus = 4
	psQuestion   PriorityStatus = 5
	psSuggestion PriorityStatus = 6
)

// BugModel the model of bug
type BugModel struct {
	ID          IndexType      `orm:"index"`      // index
	Title       string         `orm:"size(512)"`  // bug title
	Description string         `orm:"size(4096)"` // description about bug
	Version     string         `orm:"index"`      // test version number
	Source      string         // source feature request
	Target      string         `orm:"index"` // target milestone
	DevPeriod   string         `orm:"index"` // sprint
	SolveDate   time.Time      // date solving
	CreateDate  time.Time      // date creating
	Status      BugStatus      `orm:"index"` // bug current status
	Priority    PriorityStatus `orm:"index"` // bug's priority type
	Creator     IndexType      `orm:"index"` // bug's founder
	Assignor    IndexType      `orm:"index"` // who should solve this bug
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
	return nil, nil
}

// UpdateBug update bug model data
func UpdateBug(id IndexType, params map[string]interface{}) error {
	return nil
}

// DeleteBug delete bug with id
func DeleteBug(id IndexType) error {
	return nil
}
