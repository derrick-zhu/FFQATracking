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

// ContentWithIndex get content obj with its index
func (c DataPickerTemplateModel) ContentWithIndex(idx int64) interface{} {
	if idx < 0 || idx >= int64(len(c.Collection)) {
		return nil
	}

	return c.Collection[idx]
}
