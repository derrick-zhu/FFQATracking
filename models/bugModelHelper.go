package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
)

func init() {

	beego.AddFuncMap("PropertyInIssue", PropertyInIssue)
	beego.AddFuncMap("IssueCSSWithPriority", IssueCSSWithPriority)
}

// PropertyInIssue fetch issue's property value
func PropertyInIssue(pname string, value interface{}) string {

	switch value.(type) {
	case BugModel:
		issue, err := value.(BugModel)
		if err == false {
			beego.Error(err)
			return ""
		}
		_, result := BugGetReadableProperty(pname, &issue)
		return result

	default:
		beego.Error(value)
		beego.Error(errors.New("Invalid issue data type"))
	}

	return ""
}

// IntValueInIssue get int value from the BugModel
func IntValueInIssue(pname string, issue *BugModel) int64 {
	nVal, _ := BugGetReadableProperty(pname, issue)
	return nVal
}

// IssueCSSWithPriority get css style according to issue's priority level
func IssueCSSWithPriority(value interface{}) string {

	switch value.(type) {
	case BugModel:
		issue, err := value.(BugModel)
		if err == false {
			beego.Error(err)
			return ""
		}

		return fmt.Sprintf("issue-level-%d", int64(issue.Priority))

	default:
		return ""
	}
}
