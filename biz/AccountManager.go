package biz

import (
	"FFQATracking/constants"
	"FFQATracking/models"
	"FFQATracking/utils"
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// AccountManager class
type AccountManager struct {
	Account *models.AccountModel
}

// gAccMgrInstance globel AccountManager instance
var gAccMgrInstance *AccountManager
var gCookieOnce sync.Once

// AccountManagerInstance singleton
func AccountManagerInstance() *AccountManager {
	gCookieOnce.Do(func() {
		gAccMgrInstance = &AccountManager{}
	})
	return gAccMgrInstance
}

// Login login with account's email and pwd
func (am *AccountManager) Login(ctx *context.Context, email, pwd string) (bool, error) {

	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}

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

	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}

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

	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}

	ckEmail := utils.CookieInstance().Get(ctx, constants.KeyEMAIL)
	if len(ckEmail) <= 0 {
		return false
	}

	ckPwd := utils.CookieInstance().GetSecret(ctx, constants.KeyPWD)
	if len(ckPwd) <= 0 {
		return false
	}

	acc, err := am.AccountWithEMail(ckEmail)
	if err != nil {
		beego.Info("Fails in fetching account: " + ckEmail)
		beego.Error(err)
		return false
	}
	// beego.Info("Got account: " + acc.Email)

	return (acc.Pwd == utils.Base64Encode(utils.MD5(ckPwd)))
}

// CurrentAccount get current signed up account
func (am *AccountManager) CurrentAccount(ctx *context.Context) (*models.AccountModel, error) {

	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}

	if am.HadLogin(ctx) == false {
		return nil, errors.New("No user account login")
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
	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}

	return models.AccountWithID(id)
}

// AccountWithUname fetch user account with uname
func (am *AccountManager) AccountWithUname(uname string) (*models.AccountModel, error) {
	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}
	return models.AccountWithUname(uname)
}

// AccountWithEMail fetch user account with email
func (am *AccountManager) AccountWithEMail(email string) (*models.AccountModel, error) {
	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}
	return models.AccountWithEmail(email)
}

// CheckAccount check user account is matched in db
func (am *AccountManager) CheckAccount(uname, pwd string) (bool, *models.AccountModel, error) {

	if am != AccountManagerInstance() {
		log.Fatal("caller should using singleton handler")
	}

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
