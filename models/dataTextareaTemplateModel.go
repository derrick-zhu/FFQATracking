package models

// DataTextareaTemplateModel model type for textare
type DataTextareaTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue string
	Value        string
}

// ControllerType implements VarModelProtocol interface
func (c DataTextareaTemplateModel) ControllerType() TemplateDataCtrlType {
	return TextArea
}
