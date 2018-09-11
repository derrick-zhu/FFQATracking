package models

import (
	"FFQATracking/utils"
	"errors"
	"strconv"

	"github.com/astaxie/beego"
)

func init() {

	beego.AddFuncMap("PropertyInIssue", PropertyInIssue)
}

func PropertyInIssue(pname string, value interface{}) string {

	switch value.(type) {
	case BugModel:
		issue, err := value.(BugModel)
		if err == false {
			beego.Error(err)
			return ""
		}

		switch pname {
		case "ID":
			return strconv.Itoa(int(issue.ID))

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
			return strconv.FormatInt(issue.Status, 10)

		case "Priority":
			return strconv.FormatInt(issue.Priority, 10)

		case "Creator":
			return strconv.FormatInt(int64(issue.Creator), 10)

		case "Assignor":
			return strconv.FormatInt(int64(issue.Assignor), 10)
		}

	default:
		beego.Error(errors.New("Invalid issue data type"))
	}

	return ""
}
