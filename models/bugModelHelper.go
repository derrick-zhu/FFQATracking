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
		return GetReadableProperty(pname, issue)

	default:
		beego.Error(errors.New("Invalid issue data type"))
	}

	return ""
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
