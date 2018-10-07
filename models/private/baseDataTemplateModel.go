package private

type TemplateDataCtrlType int

const (
	_ TemplateDataCtrlType = iota
	TextField
	TextArea
	Number
)

// BaseDataTemplateModel basic data template struct for input
type BaseDataTemplateModel struct {
	ID         int64 // just ID for any index num
	Title      string
	Identifier string
	Type       TemplateDataCtrlType
}
