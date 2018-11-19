package models

import (
	"FFQATracking/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	// BugsTable db table name for bugmodel
	BugsTable string = "bugmodel"
)

// BugStatusModel for bug status
type BugStatusModel struct {
	dataType    int64
	description string
}

func (c BugStatusModel) Type() int64  { return c.dataType }
func (c BugStatusModel) Desc() string { return c.description }

// BugPriorityModel priority model
type BugPriorityModel struct {
	dataType    int64
	description string
}

func (c BugPriorityModel) Type() int64  { return c.dataType }
func (c BugPriorityModel) Desc() string { return c.description }

// BugReproductableModel reproductability model
type BugReproductableModel struct {
	dataType    int64
	description string
}

func (c BugReproductableModel) Type() int64  { return c.dataType }
func (c BugReproductableModel) Desc() string { return c.description }

var (
	// BugNew NEW
	BugNew = BugStatusModel{dataType: 0, description: "New"}
	// BugFixed NEW
	BugFixed = BugStatusModel{dataType: 1, description: "Fixed"}
	// BugReopen NEW
	BugReopen = BugStatusModel{dataType: 2, description: "Reopen"}
	// BugConfirm NEW
	BugConfirm = BugStatusModel{dataType: 3, description: "Confirm"}
	// BugClose NEW
	BugClose = BugStatusModel{dataType: 4, description: "Close"}
	// BugNotABug NEW
	BugNotABug = BugStatusModel{dataType: 5, description: "Not a bug"}
	// BugWillNotFix NEW
	BugWillNotFix = BugStatusModel{dataType: 6, description: "Will not fix"}
	// BugDelay NEW
	BugDelay = BugStatusModel{dataType: 7, description: "Delay"}
	// BugMustBeFix NEW
	BugMustBeFix = BugStatusModel{dataType: 8, description: "Must be fix"}
)

// AllBugStatus collections of all bug status
var AllBugStatus = []VarModelProtocol{
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
	PriorityUrgent     = BugPriorityModel{dataType: IssueUrget, description: "Urgent"}
	PriorityImportant  = BugPriorityModel{dataType: IssueImportant, description: "Important"}
	PriorityHigh       = BugPriorityModel{dataType: IssueHigh, description: "High"}
	PriorityMiddle     = BugPriorityModel{dataType: IssueMiddle, description: "Middle"}
	PriorityLow        = BugPriorityModel{dataType: IssueLow, description: "Low"}
	PriorityQuestion   = BugPriorityModel{dataType: IssueQuestion, description: "Question"}
	PrioritySuggestion = BugPriorityModel{dataType: IssueSuggestion, description: "Suggestion"}
)

// AllPriorities collections of all bug priority
var AllPriorities = []VarModelProtocol{
	PriorityUrgent,
	PriorityImportant,
	PriorityHigh,
	PriorityMiddle,
	PriorityLow,
	PriorityQuestion,
	PrioritySuggestion,
}

var (
	Reproductability100 = BugReproductableModel{dataType: 0, description: "100%"}
	Reproductability80  = BugReproductableModel{dataType: 1, description: "80%"}
	Reproductability60  = BugReproductableModel{dataType: 2, description: "60%"}
	Reproductability40  = BugReproductableModel{dataType: 3, description: "40%"}
	Reproductability20  = BugReproductableModel{dataType: 4, description: "20%"}
)

var AllReproductabilities = []VarModelProtocol{
	Reproductability100,
	Reproductability80,
	Reproductability60,
	Reproductability40,
	Reproductability20,
}

// BugModel the model of bug
type BugModel struct {
	ID               int64  `orm:"index;pk;auto"` // index
	Title            string `orm:"size(512)"`     // bug title
	Description      string `orm:"size(4096)"`    // description about bug
	FoundInSprint    int64  `orm:"index"`         // sprint - sprint
	FoundInProject   int64  `orm:"index"`         // source feature request - project
	FoundInVersion   int64  `orm:"index"`         // test version number - version
	CreateDate       int64  // date creating
	SolveDate        int64  // date solving
	LastUpdateDate   int64  // date about latest update
	Status           int64  // bug current status
	Priority         int64  // bug's priority type
	Creator          int64  // bug's founder
	Assignor         int64  // who should solve this bug
	Reproductability int64  // 重现概率 0~100
}

func init() {
	orm.RegisterModel(new(BugModel))

	beego.AddFuncMap("BugStatusWithType", BugStatusWithType)
}

// TableName for beego using
func (c *BugModel) TableName() string {
	return BugsTable
}

// EnumAllBugsStatus all bugs status collection in string
func EnumAllBugsStatus() []string {

	var result []string

	for _, eachBugModel := range AllBugStatus {
		result = append(result, eachBugModel.Desc())
	}

	return result
}

func BugStatusWithString(str string) int64 {

	for _, eachStatus := range AllBugStatus {
		if eachStatus.Desc() == str {
			return eachStatus.Type()
		}
	}
	return 0
}

func BugStatusWithType(status int64) string {
	for _, eachStatus := range AllBugStatus {
		if eachStatus.Type() == status {
			return eachStatus.Desc()
		}
	}
	return "-"
}

// EnumAllBugsPriority all priority collection in string
func EnumAllBugsPriority() []string {

	var result []string

	for _, eachPriority := range AllPriorities {
		result = append(result, eachPriority.Desc())
	}

	return result
}

func BugPriorityWithString(str string) int64 {

	for _, eachPriority := range AllPriorities {
		if eachPriority.Desc() == str {
			return eachPriority.Type()
		}
	}
	return 0
}

func BugPriorityWithType(priority int64) string {
	for _, eachPriority := range AllPriorities {
		if eachPriority.Type() == priority {
			return eachPriority.Desc()
		}
	}
	return "-"
}

func EnumAllReproductabilities() []string {

	var result []string

	for _, eachRepro := range AllReproductabilities {
		result = append(result, eachRepro.Desc())
	}

	return result
}

func BugReproductabilityWithString(str string) int64 {
	for _, eachRepro := range AllReproductabilities {
		if eachRepro.Desc() == str {
			return eachRepro.Type()
		}
	}
	return 0
}

func BugReproductabilityWithType(reproduct int64) string {
	for _, eachRepro := range AllReproductabilities {
		if eachRepro.Type() == reproduct {
			return eachRepro.Desc()
		}
	}
	return "-"
}

func BugGetReadableProperty(pname string, issue *BugModel) (int64, string) {

	switch pname {
	case "ID":
		return int64(issue.ID), strconv.FormatInt(int64(issue.ID), 10)

	case "Title":
		return 0, issue.Title

	case "Description":
		return 0, issue.Description

	case "FoundInVersion":
		return issue.FoundInVersion, ""

	case "FoundInProject":
		return issue.FoundInProject, ""

	case "FoundInSprint":
		return issue.FoundInSprint, ""

	case "SolveDate":
		return int64(issue.SolveDate), utils.StandardFormatedTimeFromTick(int64(issue.SolveDate))

	case "CreateDate":
		return int64(issue.CreateDate), utils.StandardFormatedTimeFromTick(int64(issue.CreateDate))

	case "Status":
		return issue.Status, BugStatusWithType(issue.Status)

	case "Priority":
		return issue.Priority, BugPriorityWithType(issue.Priority)

	case "Reproductability":
		return issue.Reproductability, BugReproductabilityWithType(issue.Reproductability)

	case "Creator":
		acc, err := AccountWithID(issue.Creator)
		if err != nil {
			beego.Error(err)
			return 0, "-"
		}
		return issue.Creator, acc.Name

	case "Assignor":
		acc, err := AccountWithID(issue.Assignor)
		if err != nil {
			beego.Error(err)
			return 0, "-"
		}
		return issue.Assignor, acc.Name

	default:
		return 0, "-"
	}
}

// AddBug insert new bug
func AddBug(title, description string, status, priority, creatorID, assignorID, reproductRatio int64) (*BugModel, error) {
	pBug := &BugModel{
		Title:            title,
		Description:      description,
		Status:           status,
		Priority:         priority,
		Creator:          creatorID,
		CreateDate:       utils.TimeTickSince1970(),
		Assignor:         assignorID,
		Reproductability: reproductRatio,
	}

	o, _ := GetQuerySeterWithTable(BugsTable)
	if _, err := o.Insert(pBug); err != nil {
		beego.Error(err)
		o.Rollback()

		return nil, err
	}

	return pBug, nil
}

// BugWithID fetch bug with id
func BugWithID(id int64) (*BugModel, error) {

	beego.Info("BugWithID: ", id)
	pbug := &BugModel{ID: id}

	o, _ := GetQuerySeterWithTable(BugsTable)
	if err := o.Read(pbug); err != nil {
		beego.Error(err)
		return nil, err
	}

	return pbug, nil
}

// BugsWithRange fetch bug data with range [offset, offset + count)
func BugsWithRange(offset, count int) (*[]BugModel, error) {

	var result = &[]BugModel{}
	var err error
	var rawResult orm.RawSeter

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", BugsTable, count, offset)
	rawResult = o.Raw(sqlQuery)

	if _, err = rawResult.QueryRows(result); err != nil {

		beego.Error(err)
		return nil, err
	}

	return result, nil
}

// BugsFromProjectID fetch bugs which belong some project.
func BugsFromProjectID(sprintID, projectID, versionID, offset, count int64) (*[]BugModel, error) {

	var result = &[]BugModel{}
	var sqlQuery string
	var err error
	var rawResult orm.RawSeter

	var conditionQuerySlice []string

	if sprintID >= 0 {
		conditionQuerySlice = append(conditionQuerySlice, fmt.Sprintf("found_in_sprint = %d", sprintID))
	}

	if projectID >= 0 {
		conditionQuerySlice = append(conditionQuerySlice, fmt.Sprintf("found_in_project = %d", projectID))
	}

	if versionID >= 0 {
		conditionQuerySlice = append(conditionQuerySlice, fmt.Sprintf("found_in_version = %d", versionID))
	}

	o := GetOrmObject()

	if len(conditionQuerySlice) > 0 {
		sqlQuery = fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT %d OFFSET %d;", BugsTable, strings.Join(conditionQuerySlice, " AND "), count, offset)
	} else {
		sqlQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d;", BugsTable, count, offset)
	}

	rawResult = o.Raw(sqlQuery)

	if _, err = rawResult.QueryRows(result); err != nil {

		beego.Error(err)
		return nil, err
	}

	return result, nil
}

// AllBugsData fetch all bugs
func AllBugsData() (*[]BugModel, error) {

	return BugsWithRange(0, -1)
}

// UpdateBug update bug model data
func UpdateBug(pBug *BugModel) error {

	o, _ := GetQuerySeterWithTable(BugsTable)

	if _, err := o.Update(pBug); err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}

	return nil
}

// DeleteBug delete bug with id
func DeleteBug(id int64) error {

	o, _ := GetQuerySeterWithTable(BugsTable)

	bug := &BugModel{ID: id}
	if _, err := o.Delete(bug); err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}

	return nil
}
