package models

// DataFieldTemplateModel input field template model
type DataFieldTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue string
	Value        string
}

func (c DataFieldTemplateModel) ControllerType() TemplateDataCtrlType {
	return TextField
}
