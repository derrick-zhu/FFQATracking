package utils_test

import (
	"FFQATracking/utils"
	"testing"
)

type reflectFoo struct {
	varStr string
	varInt int32
}

func TestFieldInObject(t *testing.T) {
	pFoo := &reflectFoo{"hello world", 123}
	if varStrTest := utils.FieldInObject("varStr", pFoo); varStrTest != "hello world" {
		t.Errorf("FieldInObject was incorrect, got: %s, want: %s.", varStrTest, "hello world")
	}
}
