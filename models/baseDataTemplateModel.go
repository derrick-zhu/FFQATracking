package models

import (
	"reflect"

	"github.com/astaxie/beego"
)

// TemplateDataCtrlType data controller for html components
type TemplateDataCtrlType int

const (
	_ TemplateDataCtrlType = iota
	TextField
	TextArea
	Number
)

// VarModelProtocol common variable model
type VarModelProtocol interface {
	Type() int64
	Desc() string
}

// DataControllerTypeProtocol for identify the controller type
type DataControllerTypeProtocol interface {
	ControllerType() TemplateDataCtrlType
}

// DataBaseTemplateModel basic data template struct for input
type DataBaseTemplateModel struct {
	ID         int64 // just ID for any index num
	Title      string
	Identifier string
	Type       TemplateDataCtrlType
}

func init() {
	beego.AddFuncMap("ControllerTypeOfTemplateData", ControllerTypeOfTemplateData)
}

// ControllerType DataBaseTemplateModel's implementation
func (c DataBaseTemplateModel) ControllerType() TemplateDataCtrlType {
	return TextField
}

// ControllerTypeOfTemplateData get controller type of this template model
func ControllerTypeOfTemplateData(tplData DataControllerTypeProtocol) TemplateDataCtrlType {

	typeName := reflect.TypeOf(tplData).Name()
	beego.Info(typeName, " ", tplData.ControllerType())

	return tplData.ControllerType()
}
