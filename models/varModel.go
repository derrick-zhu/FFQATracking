package models

import (
	"github.com/astaxie/beego"
)

// VarModel common variable model
type VarModel struct {
	Type int64
	Desc string
}

// BugStatusModel for bug status
type BugStatusModel struct {
	VarModel
}

// BugPriorityModel priority model
type BugPriorityModel struct {
	VarModel
}

// BugReproductableModel reproductability model
type BugReproductableModel struct {
	VarModel
}

func init() {

	beego.AddFuncMap("GetTypeFromModel", GetTypeFromModel)
	beego.AddFuncMap("GetBriefTitleFromModel", GetBriefTitleFromModel)
}

// GetTypeFromModel get Type data from value which should be VarModel
func GetTypeFromModel(value interface{}) int64 {

	switch value.(type) {
	case VarModel:
		newValue, err := value.(VarModel)
		if err == false {
			beego.Error(err)
			return 0
		}
		return newValue.Type

	case BugStatusModel:
		newValue, err := value.(BugStatusModel)
		if err == false {
			beego.Error(err)
			return 0
		}
		return newValue.Type

	case BugPriorityModel:
		newValue, err := value.(BugPriorityModel)
		if err == false {
			beego.Error(err)
			return 0
		}
		return newValue.Type

	case BugReproductableModel:
		newValue, err := value.(BugReproductableModel)
		if err == false {
			beego.Error(err)
			return 0
		}
		return newValue.Type

	case AccountModel:
		newValue, err := value.(AccountModel)
		if err == false {
			beego.Error(err)
			return 0
		}
		return int64(newValue.ID)

	default:
		beego.Info("other", value)
	}
	return 0
}

// GetBriefTitleFromModel get Desc data from value which should be VarModel
func GetBriefTitleFromModel(value interface{}) string {

	switch value.(type) {
	case VarModel:
		newValue, err := value.(VarModel)
		if err == false {
			beego.Error(err)
			return ""
		}
		return newValue.Desc

	case BugStatusModel:
		newValue, err := value.(BugStatusModel)
		if err == false {
			beego.Error(err)
			return ""
		}
		return newValue.Desc

	case BugPriorityModel:
		newValue, err := value.(BugPriorityModel)
		if err == false {
			beego.Error(err)
			return ""
		}
		return newValue.Desc

	case BugReproductableModel:
		newValue, err := value.(BugReproductableModel)
		if err == false {
			beego.Error(err)
			return ""
		}
		return newValue.Desc

	case AccountModel:
		newValue, err := value.(AccountModel)
		if err == false {
			beego.Error(err)
			return ""
		}
		return newValue.Name

	default:
		beego.Info("other", value)
	}
	return ""
}