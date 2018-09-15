package biz

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"FFQATracking/utils"
	"errors"
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// AccountManager class
type AccountManager struct {
	Account *models.AccountModel
}

// AccMgrInstance globel AccountManager instance
var AccMgrInstance *AccountManager
var cookieOnce sync.Once

// AccountManagerInstance singleton
func AccountManagerInstance() *AccountManager {
	cookieOnce.Do(func() {
		AccMgrInstance = &AccountManager{}
	})
	return AccMgrInstance
}

// Login login with account's email and pwd
func (am *AccountManager) Login(ctx *context.Context, email, pwd string) (bool, error) {

	result, acc, _ := AccountManagerInstance().CheckAccount(email, pwd)
	if true == result {

		utils.CookieInstance().Set(ctx, constants.KeyUID, strconv.Itoa(int(acc.ID)), -1)
		utils.CookieInstance().Set(ctx, constants.KeyEMAIL, email, -1)
		utils.CookieInstance().SetSecret(ctx, constants.KeyPWD, pwd, -1)

		return true, nil
	}

	return false, errors.New("invalid user account or password")
}

// Logout logout
func (am *AccountManager) Logout(ctx *context.Context, uid string) bool {

	var err error
	var id int64

	id, err = strconv.ParseInt(uid, 10, 64)
	if err != nil {
		beego.Error(err)
		return false
	}

	_, err = AccountManagerInstance().AccountWithID(id)
	if err != nil {
		beego.Error("user account not existed.")
		return false
	}

	utils.CookieInstance().Set(ctx, constants.KeyUID, "", -1)
	utils.CookieInstance().Set(ctx, constants.KeyEMAIL, "", -1)
	utils.CookieInstance().SetSecret(ctx, constants.KeyPWD, "", -1)

	return true
}

// HadLogin check account login state
func (am *AccountManager) HadLogin(ctx *context.Context) bool {

	ckEmail := utils.CookieInstance().Get(ctx, constants.KeyEMAIL)
	beego.Info("ckEmail = " + ckEmail)
	if len(ckEmail) <= 0 {
		return false
	}

	ckPwd := utils.CookieInstance().GetSecret(ctx, constants.KeyPWD)
	beego.Info("ckPwd = " + ckPwd)
	if len(ckPwd) <= 0 {
		return false
	}

	acc, err := am.AccountWithEMail(ckEmail)
	if err != nil {
		beego.Info("Fails in fetching account: " + ckEmail)
		beego.Error(err)
		return false
	}
	beego.Info("Got account: " + acc.Email)

	return (acc.Pwd == utils.Base64Encode(utils.MD5(ckPwd)))
}

// CurrentAccount get current signed up account
func (am *AccountManager) CurrentAccount(ctx *context.Context) (*models.AccountModel, error) {

	if am.HadLogin(ctx) == false {
		return nil, errors.New("Not login")
	}

	ckEmail := utils.CookieInstance().Get(ctx, constants.KeyEMAIL)
	acc, err := am.AccountWithEMail(ckEmail)
	if err != nil {
		beego.Info("Fails in fetching account: " + ckEmail)
		beego.Error(err)

		return nil, err
	}

	return acc, nil
}

// AccountWithID fetch user account with uid
func (am *AccountManager) AccountWithID(id int64) (*models.AccountModel, error) {
	return models.AccountWithID(id)
}

// AccountWithUname fetch user account with uname
func (am *AccountManager) AccountWithUname(uname string) (*models.AccountModel, error) {
	return models.AccountWithUname(uname)
}

// AccountWithEMail fetch user account with email
func (am *AccountManager) AccountWithEMail(email string) (*models.AccountModel, error) {
	return models.AccountWithEmail(email)
}

// CheckAccount check user account is matched in db
func (am *AccountManager) CheckAccount(uname, pwd string) (bool, *models.AccountModel, error) {

	var acc *models.AccountModel
	var err error

	if utils.MatchRegexEmail(uname) {
		acc, err = models.AccountWithEmail(uname)
	} else {
		acc, err = models.AccountWithUname(uname)
	}

	if err != nil {
		beego.Error(err)
		return false, nil, err
	}

	base64Pwd := utils.Base64Encode(utils.MD5(pwd))
	return (acc.Pwd == base64Pwd), acc, nil
}
