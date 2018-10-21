package models

// DataDatePickerTemplateModel picker's model
type DataDatePickerTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue int64
	Value        int64
}

// ControllerType implement VarModelProtocol
func (c DataDatePickerTemplateModel) ControllerType() TemplateDataCtrlType {
	return Date
}
