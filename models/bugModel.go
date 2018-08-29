package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// BugStatus bug status type
type BugStatus string

const (
	bsNew        BugStatus = "New"
	bsFixed      BugStatus = "Fixed"
	bsReopen     BugStatus = "Reopen"
	bsConfirm    BugStatus = "Confirm"
	bsClose      BugStatus = "Close"
	bsNotABug    BugStatus = "Not a bug"
	bsWillNotFix BugStatus = "Will not fix"
	bsDelay      BugStatus = "Delay"
	bsMustBeFix  BugStatus = "Must be fix"
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

func (this *BugModel) TableName() string {
	return "bugmodel"
}
