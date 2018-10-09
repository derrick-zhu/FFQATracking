package models

// DataPickerTemplateModel class template
type DataPickerTemplateModel struct {
	BaseDataTemplateModel

	DefaultValue int64
	Value        int64
	Collection   []VarModelProtocol
}

func (c DataPickerTemplateModel) ControllerType() TemplateDataCtrlType {
	return Number
}
