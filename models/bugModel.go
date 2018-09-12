package models

import (
	"FFQATracking/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	// BugsTable db table name for bugmodel
	BugsTable string = "bugmodel"
)

var (
	// BugNew NEW
	BugNew = BugStatusModel{VarModel: VarModel{Type: 0, Desc: "New"}}
	// BugFixed NEW
	BugFixed = BugStatusModel{VarModel: VarModel{Type: 1, Desc: "Fixed"}}
	// BugReopen NEW
	BugReopen = BugStatusModel{VarModel: VarModel{Type: 2, Desc: "Reopen"}}
	// BugConfirm NEW
	BugConfirm = BugStatusModel{VarModel: VarModel{Type: 3, Desc: "Confirm"}}
	// BugClose NEW
	BugClose = BugStatusModel{VarModel: VarModel{Type: 4, Desc: "Close"}}
	// BugNotABug NEW
	BugNotABug = BugStatusModel{VarModel: VarModel{Type: 5, Desc: "Not a bug"}}
	// BugWillNotFix NEW
	BugWillNotFix = BugStatusModel{VarModel: VarModel{Type: 6, Desc: "Will not fix"}}
	// BugDelay NEW
	BugDelay = BugStatusModel{VarModel: VarModel{Type: 7, Desc: "Delay"}}
	// BugMustBeFix NEW
	BugMustBeFix = BugStatusModel{VarModel: VarModel{Type: 8, Desc: "Must be fix"}}
)

// AllBugStatus collections of all bug status
var AllBugStatus = []BugStatusModel{
	BugNew,
	BugFixed,
	BugReopen,
	BugConfirm,
	BugClose,
	BugNotABug,
	BugWillNotFix,
	BugDelay,
	BugMustBeFix,
}

const (
	IssueUrget      int64 = 0
	IssueImportant  int64 = 1
	IssueHigh       int64 = 2
	IssueMiddle     int64 = 3
	IssueLow        int64 = 4
	IssueQuestion   int64 = 5
	IssueSuggestion int64 = 6
)

var (
	PriorityUrgent     = BugPriorityModel{VarModel: VarModel{Type: IssueUrget, Desc: "Urgent"}}
	PriorityImportant  = BugPriorityModel{VarModel: VarModel{Type: IssueImportant, Desc: "Important"}}
	PriorityHigh       = BugPriorityModel{VarModel: VarModel{Type: IssueHigh, Desc: "High"}}
	PriorityMiddle     = BugPriorityModel{VarModel: VarModel{Type: IssueMiddle, Desc: "Middle"}}
	PriorityLow        = BugPriorityModel{VarModel: VarModel{Type: IssueLow, Desc: "Low"}}
	PriorityQuestion   = BugPriorityModel{VarModel: VarModel{Type: IssueQuestion, Desc: "Question"}}
	PrioritySuggestion = BugPriorityModel{VarModel: VarModel{Type: IssueSuggestion, Desc: "Suggestion"}}
)

// AllPriorities collections of all bug priority
var AllPriorities = []BugPriorityModel{
	PriorityUrgent,
	PriorityImportant,
	PriorityHigh,
	PriorityMiddle,
	PriorityLow,
	PriorityQuestion,
	PrioritySuggestion,
}

var (
	Reproductability100 = BugReproductableModel{VarModel: VarModel{Type: 0, Desc: "100%"}}
	Reproductability80  = BugReproductableModel{VarModel: VarModel{Type: 1, Desc: "80%"}}
	Reproductability60  = BugReproductableModel{VarModel: VarModel{Type: 2, Desc: "60%"}}
	Reproductability40  = BugReproductableModel{VarModel: VarModel{Type: 3, Desc: "40%"}}
	Reproductability20  = BugReproductableModel{VarModel: VarModel{Type: 4, Desc: "20%"}}
)

var AllReproductabilities = []BugReproductableModel{
	Reproductability100,
	Reproductability80,
	Reproductability60,
	Reproductability40,
	Reproductability20,
}

// BugModel the model of bug
type BugModel struct {
	ID               IndexType    `orm:"index"`      // index
	Title            string       `orm:"size(512)"`  // bug title
	Description      string       `orm:"size(4096)"` // description about bug
	Version          string       `orm:"index"`      // test version number
	Source           string       // source feature request
	Target           string       `orm:"index"` // target milestone
	DevPeriod        string       `orm:"index"` // sprint
	SolveDate        TimeInterval // date solving
	CreateDate       TimeInterval // date creating
	Status           int64        `orm:"index"` // bug current status
	Priority         int64        `orm:"index"` // bug's priority type
	Creator          IndexType    `orm:"index"` // bug's founder
	Assignor         IndexType    `orm:"index"` // who should solve this bug
	Reproductability int64        // 重现概率 0~100
}

func init() {
	orm.RegisterModel(new(BugModel))
}

// TableName for beego using
func (c *BugModel) TableName() string {
	return BugsTable
}

// EnumAllBugsStatus all bugs status collection in string
func EnumAllBugsStatus() []string {

	var result []string

	for _, eachBugModel := range AllBugStatus {
		result = append(result, eachBugModel.Desc)
	}

	return result
}

func BugStatusWithString(str string) int64 {

	for _, eachStatus := range AllBugStatus {
		if eachStatus.Desc == str {
			return eachStatus.Type
		}
	}
	return 0
}

func BugStatusWithType(status int64) string {
	for _, eachStatus := range AllBugStatus {
		if eachStatus.Type == status {
			return eachStatus.Desc
		}
	}
	return "-"
}

// EnumAllBugsPriority all priority collection in string
func EnumAllBugsPriority() []string {

	var result []string

	for _, eachPriority := range AllPriorities {
		result = append(result, eachPriority.Desc)
	}

	return result
}

func BugPriorityWithString(str string) int64 {

	for _, eachPriority := range AllPriorities {
		if eachPriority.Desc == str {
			return eachPriority.Type
		}
	}
	return 0
}

func BugPriorityWithType(priority int64) string {
	for _, eachPriority := range AllPriorities {
		if eachPriority.Type == priority {
			return eachPriority.Desc
		}
	}
	return "-"
}

func EnumAllReproductabilities() []string {

	var result []string

	for _, eachRepro := range AllReproductabilities {
		result = append(result, eachRepro.Desc)
	}

	return result
}

func BugReproductabilityWithString(str string) int64 {
	for _, eachRepro := range AllReproductabilities {
		if eachRepro.Desc == str {
			return eachRepro.Type
		}
	}
	return 0
}

func BugReproductabilityWithType(reproduct int64) string {
	for _, eachRepro := range AllReproductabilities {
		if eachRepro.Type == reproduct {
			return eachRepro.Desc
		}
	}
	return "-"
}

func GetReadableProperty(pname string, issue BugModel) string {

	switch pname {
	case "ID":
		return strconv.FormatInt(int64(issue.ID), 10)

	case "Title":
		return issue.Title

	case "Description":
		return issue.Description

	case "Version":
		return issue.Version

	case "Source":
		return issue.Source

	case "Target":
		return issue.Target

	case "DevPeriod":
		return issue.DevPeriod

	case "SolveDate":
		return utils.StandardFormatedTimeFromTick(int64(issue.SolveDate))

	case "CreateDate":
		return utils.StandardFormatedTimeFromTick(int64(issue.CreateDate))

	case "Status":
		return BugStatusWithType(issue.Status)

	case "Priority":
		return BugPriorityWithType(issue.Priority)

	case "Reproductability":
		return BugReproductabilityWithType(issue.Reproductability)

	case "Creator":
		acc, err := AccountWithID(issue.Creator)
		if err != nil {
			beego.Error(err)
			return "-"
		}
		return acc.Name

	case "Assignor":
		acc, err := AccountWithID(issue.Assignor)
		if err != nil {
			beego.Error(err)
			return "-"
		}
		return acc.Name

	default:
		return "-"
	}
}

// AddBug insert new bug
func AddBug(title, description string, status, priority, creatorID, assignorID, reproductRatio int64) (*BugModel, error) {
	pBug := &BugModel{
		Title:            title,
		Description:      description,
		Status:           status,
		Priority:         priority,
		Creator:          IndexType(creatorID),
		CreateDate:       TimeInterval(utils.TimeIntervalSince1970()),
		Assignor:         IndexType(assignorID),
		Reproductability: reproductRatio,
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
func BugsWithRange(lower, count int) ([]BugModel, error) {

	var result []BugModel
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

// AllBugsData fetch all bugs
func AllBugsData() ([]BugModel, error) {

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
