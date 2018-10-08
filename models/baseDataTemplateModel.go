package models

import (
	"github.com/astaxie/beego"
)

type TemplateDataCtrlType int

const (
	_ TemplateDataCtrlType = iota
	TextField
	TextArea
	Number
)

type DataControllerTypeProtocol interface {
	ControllerType() TemplateDataCtrlType
}

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

func (c *BaseDataTemplateModel) DataControllerType() TemplateDataCtrlType {
	return TextField
}

// ControllerTypeOfTemplateData get controller type of this template model
func ControllerTypeOfTemplateData(tplData DataControllerTypeProtocol) TemplateDataCtrlType {

	// beego.Info(tplData)
	// typeDefine := reflect.TypeOf(tplData)
	// typeName := typeDefine.Name()
	// beego.Info(typeName)

	return tplData.ControllerType()
}
