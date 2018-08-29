package biz

import (
	"FFQATracking/models"
	"FFQATracking/utils"

	"github.com/astaxie/beego"
)

func CheckAccount(uname, pwd string) bool {
	acc, err := models.AccountWithUname(uname)
	if err != nil {
		beego.Error(err)
		return false
	}

	digistPwd := utils.MD5(pwd)
	return (acc.Pwd == digistPwd)
}
