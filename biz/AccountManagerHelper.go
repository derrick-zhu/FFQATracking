package biz

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"FFQATracking/utils"
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// HasAccountIfNot check whether account is existed or not
func HasAccountIfNot(uname string) bool {

	var err error

	if utils.MatchRegexEmail(uname) {
		_, err = models.AccountWithEmail(uname)
	} else {
		_, err = models.AccountWithUname(uname)
	}

	if err != nil {
		beego.Debug(err)
		return false
	}
	return true
}

// CheckAccount check user account is matched in db
func CheckAccount(email, pwd string) (bool, *models.AccountModel) {

	result, acc, _ := AccountManagerInstance().CheckAccount(email, pwd)
	return result, acc
}

// Register for helping user to register a new account
func Register(email, pwd string, rule models.RuleType) (bool, *models.AccountModel, error) {

	var nickName string

	if utils.MatchRegexEmail(email) == false {

		return false, nil, errors.New("invalid email formation")
	}

	nickName = email[0:strings.Index(email, "@")]

	acc, err := models.AddAccount(nickName, email)
	if err != nil {

		beego.Debug(err)
		return false, nil, err
	}

	acc.Rule = rule
	acc.Job = models.JobDeveloper
	acc.Pwd = utils.Base64Encode(utils.MD5(pwd))

	err = models.UpdateAccount(acc.ID, map[string]interface{}{

		"Rule": acc.Rule,
		"Job":  acc.Job,
		"Pwd":  acc.Pwd,
	})

	if err != nil {

		return false, nil, err
	}

	return true, acc, nil
}

// Login login with user account and password
func Login(ctx *context.Context, email, pwd string) (bool, error) {

	return AccountManagerInstance().Login(ctx, email, pwd)
}

// Logout logout user account with uid
func Logout(ctx *context.Context) bool {

	uid := utils.CookieInstance().Get(ctx, constants.KeyUID)
	return AccountManagerInstance().Logout(ctx, uid)
}

// HadLogin check account login state
func HadLogin(ctx *context.Context) bool {

	return AccountManagerInstance().HadLogin(ctx)
}

// CurrentAccount fetch current account
func CurrentAccount(ctx *context.Context) (*models.AccountModel, error) {

	// beego.Info("fetch current account had been signed in")
	return AccountManagerInstance().CurrentAccount(ctx)
}
