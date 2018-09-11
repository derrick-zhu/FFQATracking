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

var (
	// BugNew NEW
	BugNew BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 0, Desc: "New"}}
	// BugFixed NEW
	BugFixed BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 1, Desc: "Fixed"}}
	// BugReopen NEW
	BugReopen BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 2, Desc: "Reopen"}}
	// BugConfirm NEW
	BugConfirm BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 3, Desc: "Confirm"}}
	// BugClose NEW
	BugClose BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 4, Desc: "Close"}}
	// BugNotABug NEW
	BugNotABug BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 5, Desc: "Not a bug"}}
	// BugWillNotFix NEW
	BugWillNotFix BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 6, Desc: "Will not fix"}}
	// BugDelay NEW
	BugDelay BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 7, Desc: "Delay"}}
	// BugMustBeFix NEW
	BugMustBeFix BugStatusModel = BugStatusModel{VarModel: VarModel{Type: 8, Desc: "Must be fix"}}
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

var (
	PriorityUrgent     BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 0, Desc: "Urgent"}}
	PriorityImportant  BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 1, Desc: "Important"}}
	PriorityHigh       BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 2, Desc: "High"}}
	PriorityMiddle     BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 3, Desc: "Middle"}}
	PriorityLow        BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 4, Desc: "Low"}}
	PriorityQuestion   BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 5, Desc: "Question"}}
	PrioritySuggestion BugPriorityModel = BugPriorityModel{VarModel: VarModel{Type: 6, Desc: "Suggestion"}}
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
	Reproductability100 BugReproductableModel = BugReproductableModel{VarModel: VarModel{Type: 0, Desc: "100%"}}
	Reproductability80  BugReproductableModel = BugReproductableModel{VarModel: VarModel{Type: 1, Desc: "80%"}}
	Reproductability60  BugReproductableModel = BugReproductableModel{VarModel: VarModel{Type: 2, Desc: "60%"}}
	Reproductability40  BugReproductableModel = BugReproductableModel{VarModel: VarModel{Type: 3, Desc: "40%"}}
	Reproductability20  BugReproductableModel = BugReproductableModel{VarModel: VarModel{Type: 4, Desc: "20%"}}
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
	ID              IndexType `orm:"index"`      // index
	Title           string    `orm:"size(512)"`  // bug title
	Description     string    `orm:"size(4096)"` // description about bug
	Version         string    `orm:"index"`      // test version number
	Source          string    // source feature request
	Target          string    `orm:"index"` // target milestone
	DevPeriod       string    `orm:"index"` // sprint
	SolveDate       time.Time // date solving
	CreateDate      time.Time // date creating
	Status          int64     `orm:"index"` // bug current status
	Priority        int64     `orm:"index"` // bug's priority type
	Creator         IndexType `orm:"index"` // bug's founder
	Assignor        IndexType `orm:"index"` // who should solve this bug
	Reproducibility int64     // 重现概率 0~100
}

func init() {
	orm.RegisterModel(new(BugModel))

	beego.AddFuncMap("VarModelGetType", VarModelGetType)
	beego.AddFuncMap("VarModelGetDesc", VarModelGetDesc)
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

// AddBug insert new bug
func AddBug(title, description string, creatorID IndexType) (*BugModel, error) {
	pBug := &BugModel{
		Title:       title,
		Description: description,
		Creator:     creatorID,
		CreateDate:  time.Now(),
		Status:      BugNew.Type,
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

// AllBugsData fetch all bugs
func AllBugsData() ([]*BugModel, error) {

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
