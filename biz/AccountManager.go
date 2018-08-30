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

// Login login with account's uname and pwd
func (am *AccountManager) Login(ctx *context.Context, uname, pwd string) (bool, error) {

	result, acc, _ := AccountManagerInstance().CheckAccount(uname, pwd)
	if false == result {

		utils.CookieInstance().Set(ctx, constants.KeyUID, strconv.Itoa(int(acc.ID)), -1)
		utils.CookieInstance().Set(ctx, constants.KeyUNAME, uname, -1)
		utils.CookieInstance().SetSecret(ctx, constants.KeyPWD, pwd, -1)

		return false, errors.New("invalid user account or password")
	}
	return true, nil
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
	utils.CookieInstance().Set(ctx, constants.KeyUNAME, "", -1)
	utils.CookieInstance().SetSecret(ctx, constants.KeyPWD, "", -1)

	return true
}

// HadLogin check account login state
func (am *AccountManager) HadLogin(ctx *context.Context) bool {

	ckUname := utils.CookieInstance().Get(ctx, constants.KeyUNAME)
	beego.Info("ckUname = " + ckUname)
	if len(ckUname) <= 0 {
		return false
	}

	ckPwd := utils.CookieInstance().GetSecret(ctx, constants.KeyPWD)
	beego.Info("ckPwd = " + ckPwd)
	if len(ckPwd) <= 0 {
		return false
	}

	acc, err := am.AccountWithUname(ckUname)
	if err != nil {
		beego.Error(err)
		return false
	}
	beego.Info("account = " + acc.Pwd)

	return (acc.Pwd == utils.Base64Encode(utils.MD5(ckPwd)))
}

// AccountWithID fetch user account with uid
func (am *AccountManager) AccountWithID(id int64) (*models.AccountModel, error) {
	return models.AccountWithID(models.IndexType(id))
}

// AccountWithUname fetch user account with uname
func (am *AccountManager) AccountWithUname(uname string) (*models.AccountModel, error) {
	return models.AccountWithUname(uname)
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

// CheckAccount check user account is matched in db
func CheckAccount(uname, pwd string) (bool, *models.AccountModel) {
	result, acc, _ := AccountManagerInstance().CheckAccount(uname, pwd)
	return result, acc
}

// Login login with user account and password
func Login(ctx *context.Context, uname, pwd string) (bool, error) {
	return AccountManagerInstance().Login(ctx, uname, pwd)
}

// Logout logout user account with uid
func Logout(ctx *context.Context) bool {
	uid := utils.CookieInstance().Get(ctx, constants.KeyUID)
	return AccountManagerInstance().Logout(ctx, uid)
}

// HadLogin check account login state
func HadLogin(ctx *context.Context) bool {
	beego.Info("biz.HadLogin()")
	return AccountManagerInstance().HadLogin(ctx)
}
