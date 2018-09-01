package biz

import (
	"testing"

	"github.com/astaxie/beego"
)

func TestRegister(t *testing.T) {
	_, _, err := Register("abc@abc.com", "admin", 0)
	if err != nil {
		beego.Error(err)
		t.Fatal("not implemented yet")
	}
}
