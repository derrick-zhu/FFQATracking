package models

// TemplateDataCtrlType data controller for html components
type TemplateDataCtrlType int

const (
	_ TemplateDataCtrlType = iota
	TextField
	TextArea
	Number
	Date
)

// VarModelProtocol common variable model
type VarModelProtocol interface {
	Type() int64
	Desc() string
}
