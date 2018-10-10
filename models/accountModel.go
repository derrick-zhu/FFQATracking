package models

import (
	"FFQATracking/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	// AccountTable account db name
	AccountTable string = "accountmodel"
)

// RuleType about account rule in FFQATracking platform
type RuleType int64

const (
	// RuleAdmin as admin
	RuleAdmin RuleType = 0
	// RuleUser as normal user
	RuleUser RuleType = 1
	// RuleGuest as anonymouse user
	RuleGuest RuleType = 2
)

// JobType about job type in FF
type JobType int64

const (
	// JobManager ...
	JobManager JobType = 0
	// JobLeader ...
	JobLeader JobType = 1
	// JobDeveloper ...
	JobDeveloper JobType = 2
	// JobQATester ...
	JobQATester JobType = 3
	// JobProductOwner ...
	JobProductOwner JobType = 4
)

// AccountModel user account model
type AccountModel struct {
	ID     int64  `orm:"pk;auto"` // readonly
	Name   string `orm:"index;size(128)"`
	Avatar string `orm:"null;size(1024)"`
	Email  string `orm:"index;size(1024)"` // readonly
	Create int64  // readonly
	Rule   RuleType
	Job    JobType
	Pwd    string
}

// Type implements the VarModelProtocol interface.
func (c AccountModel) Type() int64 { return c.ID }

// Desc implements the VarModelProtocol interface
func (c AccountModel) Desc() string { return c.Name }

func init() {
	orm.RegisterModel(new(AccountModel))
}

// TableName for beego using
func (c *AccountModel) TableName() string {
	return AccountTable
}

// InstallAdminAccount install the admin account, ONLY ONCE
func InstallAdminAccount() {
	acc, err := AddAccount("admin", "root@farfetch.com")
	if err != nil {
		beego.Error(err)
		return
	}

	acc.Rule = RuleAdmin
	acc.Job = JobManager
	acc.Pwd = utils.Base64Encode(utils.MD5("admin"))

	beego.Info(acc)
	if err = UpdateAccount(acc); err != nil {
		beego.Error(err)
	}
}

// AddAccount insert new account with name and email
func AddAccount(name string, email string) (*AccountModel, error) {

	account := &AccountModel{Name: name,
		Email:  email,
		Create: utils.TimeTickSince1970(),
	}

	o, qs := GetQuerySeterWithTable(AccountTable)

	filterErr := qs.Filter("email", email).One(account)
	if filterErr == nil { // account has already existed
		return account, nil
	}

	_, insertErr := o.Insert(account)
	if insertErr != nil {
		o.Rollback()
		return nil, insertErr
	}

	return account, nil
}

// AccountWithUname get account with uname
func AccountWithUname(uname string) (*AccountModel, error) {

	acc := &AccountModel{Name: uname}

	_, qs := GetQuerySeterWithTable(AccountTable)
	filterErr := qs.Filter("name", uname).One(acc)
	if filterErr != nil {
		return nil, filterErr
	}

	return acc, nil
}

// AccountWithEmail get account with email
func AccountWithEmail(email string) (*AccountModel, error) {

	acc := &AccountModel{Email: email}

	_, qs := GetQuerySeterWithTable(AccountTable)
	filterErr := qs.Filter("email", email).One(acc)
	if filterErr != nil {
		return nil, filterErr
	}

	return acc, nil
}

// AccountWithID get account with id
func AccountWithID(id int64) (*AccountModel, error) {

	acc := &AccountModel{ID: id}

	_, qs := GetQuerySeterWithTable(AccountTable)
	filterErr := qs.Filter("id", fmt.Sprintf("%d", id)).One(acc)
	if filterErr != nil {

		return nil, filterErr
	}

	return acc, nil
}

// AccountsWithRange fetch account data with rage [lower, count)
func AccountsWithRange(lower, count int) (*[]AccountModel, error) {
	var result = &[]AccountModel{}
	var err error
	var rawResult orm.RawSeter

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", AccountTable, count, lower)
	rawResult = o.Raw(sqlQuery)

	_, err = rawResult.QueryRows(result)
	if err != nil {

		beego.Error(err)
		return nil, err
	}

	return result, nil
}

// AllAccounts fetch all account
func AllAccounts() (*[]AccountModel, error) {

	return AccountsWithRange(0, -1)
}

// UpdateAccount [WIP] update account's content
func UpdateAccount(newAcc *AccountModel) error {

	var err error

	o, _ := GetQuerySeterWithTable(AccountTable)

	if _, err := AccountWithID(newAcc.ID); err != nil {
		beego.Debug("could not find account: %d.", newAcc.ID)
		return err
	}

	if _, err = o.Update(newAcc); err != nil {
		beego.Error(err)
		o.Rollback()

		return err
	}

	return nil
}

// DeleteAccount [DONE] delete account with id
func DeleteAccount(id int64) error {

	o, _ := GetQuerySeterWithTable(AccountTable)

	acc := &AccountModel{ID: id}
	_, err := o.Delete(acc)

	return err
}
