package models

// DataPickerTemplateModel class template
type DataPickerTemplateModel struct {
	DataBaseTemplateModel

	ValueChanged JSCommandModel // the JS callback when picker's value changed
	DefaultValue int64
	Value        int64
	Collection   []VarModelProtocol
}

// ControllerType implements VarModelProtocol interface
func (c DataPickerTemplateModel) ControllerType() TemplateDataCtrlType {
	return Number
}
