package models

import (
	"FFQATracking/models/private"
)

// DataPickerTemplateModel class template
type DataPickerTemplateModel struct {
	private.BaseDataTemplateModel

	DefaultValue int64
	Value        int64
	Collection   interface{}
}
