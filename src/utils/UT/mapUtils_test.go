package UT_test

import (
	"testing"
)

type fooMapUtils struct {
	VarInt int
	VarStr string
}

func TestIsEmpty(t *testing.T) {
	aFoo := fooMapUtils{VarInt: 0}

	if IsEmpty(aFoo) == true {
		t.Errorf("UT: Fails in testing IsEmpty(aFoo)")
	}

	if IsEmpty(aFoo.VarInt) == true {
		t.Errorf("UT: Fails in testing IsEmpty(aFoo.VarInt)")
	}

}
