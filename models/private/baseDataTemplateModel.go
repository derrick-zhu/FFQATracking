package private

import (
	"reflect"

	"github.com/astaxie/beego"
)

type TemplateDataCtrlType int

const (
	_ TemplateDataCtrlType = iota
	TextField
	TextArea
	Number
)

// BaseDataTemplateModel basic data template struct for input
type BaseDataTemplateModel struct {
	ID         int64 // just ID for any index num
	Title      string
	Identifier string
	Type       TemplateDataCtrlType
}

func init() {
	beego.AddFuncMap("ControllerTypeOfTemplateData", ControllerTypeOfTemplateData)
}

// ControllerTypeOfTemplateData get controller type of this template model
func ControllerTypeOfTemplateData(tplData interface{}) TemplateDataCtrlType {
	beego.Info(tplData)
	typeName := reflect.TypeOf(tplData).Name()
	beego.Info(typeName)

	return TextField
}
