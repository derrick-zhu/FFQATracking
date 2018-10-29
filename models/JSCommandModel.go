package models

// JSCommandModel for connecting js and go
type JSCommandModel struct {
	ID    string
	Name  string
	Param []string
}

type GOCommandModel struct {
	ID    string
	Name  string
	Param map[string]interface{}
}
