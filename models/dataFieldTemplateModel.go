package models

// DataFieldTemplateModel input field template model
type DataFieldTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue string
	Value        string
}

// ControllerType implements VarModelProtocol interface
func (c DataFieldTemplateModel) ControllerType() TemplateDataCtrlType {
	return TextField
}
