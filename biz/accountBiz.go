package biz

import (
	"FFQATracking/models"
	"FFQATracking/utils"

	"github.com/astaxie/beego"
)

func CheckAccount(uname, pwd string) bool {

	var acc *models.AccountModel
	var err error

	if utils.MatchRegexEmail(uname) {
		acc, err = models.AccountWithEmail(uname)
	} else {
		acc, err = models.AccountWithUname(uname)
	}

	if err != nil {
		beego.Error(err)
		return false
	}

	base64Pwd := utils.Base64Encode(utils.MD5(pwd))
	return (acc.Pwd == base64Pwd)
}
