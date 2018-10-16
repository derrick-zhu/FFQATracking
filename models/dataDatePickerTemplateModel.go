package models

type DataDatePickerTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue int64
	Value        int64
}

func (c DataDatePickerTemplateModel) ControllerType() TemplateDataCtrlType {
	return Date
}
