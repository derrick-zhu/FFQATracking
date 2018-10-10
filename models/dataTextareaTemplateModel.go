package models

// DataTextareaTemplateModel model type for textare
type DataTextareaTemplateModel struct {
	DataBaseTemplateModel

	DefaultValue string
	Value        string
}

func (c DataTextareaTemplateModel) ControllerType() TemplateDataCtrlType {
	return TextArea
}
