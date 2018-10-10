package models

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.AddFuncMap("GetTypeFromModel", GetTypeFromModel)
	beego.AddFuncMap("GetBriefTitleFromModel", GetBriefTitleFromModel)
}

// GetTypeFromModel get Type data from value which should be VarModelProtocol
func GetTypeFromModel(value interface{}) int64 {

	switch value.(type) {
	case VarModelProtocol:
		var newValue VarModelProtocol
		var ok bool
		if newValue, ok = value.(VarModelProtocol); ok != true {
			beego.Error(ok)
			return 0
		}
		return newValue.Type()

	default:
		beego.Info("other", value)
	}
	return 0
}

// GetBriefTitleFromModel get Desc data from value which should be VarModelProtocol
func GetBriefTitleFromModel(value interface{}) string {

	switch value.(type) {
	case VarModelProtocol:
		newValue, err := value.(VarModelProtocol)
		if err == false {
			beego.Error(err)
			return ""
		}
		return newValue.Desc()
		
	default:
		beego.Info("other", value)
	}
	return ""
}
