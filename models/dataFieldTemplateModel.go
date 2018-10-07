package models

import "FFQATracking/models/private"

// DataFieldTemplateModel input field template model
type DataFieldTemplateModel struct {
	private.BaseDataTemplateModel

	DefaultValue string
	Value        string
}
