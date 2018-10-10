package models

// DataPickerTemplateModel class template
type DataPickerTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue int64
	Value        int64
	Collection   []VarModelProtocol
}

// ControllerType implements VarModelProtocol interface
func (c DataPickerTemplateModel) ControllerType() TemplateDataCtrlType {
	return Number
}
