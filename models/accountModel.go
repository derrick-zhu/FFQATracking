package models

import (
	"FFQATracking/utils"
	"fmt"
	"time"

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
	RuleAdmin RuleType = 0
	RuleUser  RuleType = 1
	RuleGuest RuleType = 2
)

// JobType about job type in FF
type JobType int64

const (
	JobManager      JobType = 0
	JobLeader       JobType = 1
	JobDeveloper    JobType = 2
	JobQATester     JobType = 3
	JobProductOwner JobType = 4
)

// AccountModel user account model
type AccountModel struct {
	ID     IndexType `orm:"pk;auto"` // readonly
	Name   string    `orm:"index;size(128)"`
	Avatar string    `orm:"null;size(1024)"`
	Email  string    `orm:"index;size(1024)"` // readonly
	Create time.Time `orm:"index"`            // readonly
	Rule   RuleType
	Job    JobType
	Pwd    string
}

func init() {
	orm.RegisterModel(new(AccountModel))
}

// TableName for beego using
func (this *AccountModel) TableName() string {
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
	err = UpdateAccount(acc.ID, map[string]interface{}{
		"Rule": acc.Rule,
		"Job":  acc.Job,
		"Pwd":  acc.Pwd,
	})

	if err != nil {
		beego.Debug(err)
	}
}

// AddAccount insert new account with name and email
func AddAccount(name string, email string) (*AccountModel, error) {

	account := &AccountModel{Name: name,
		Email:  email,
		Create: time.Now(),
	}

	o, qs := GetQuerySeterWithTable(AccountTable)

	filterErr := qs.Filter("email", email).One(account)
	if filterErr == nil { // account has already existed
		return account, nil
	}

	_, insertErr := o.Insert(account)
	if insertErr != nil {
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
func AccountWithID(id IndexType) (*AccountModel, error) {

	acc := &AccountModel{ID: id}

	_, qs := GetQuerySeterWithTable(AccountTable)
	filterErr := qs.Filter("id", fmt.Sprintf("%d", id)).One(acc)
	if filterErr != nil {

		return nil, filterErr
	}

	return acc, nil
}

// AccountsWithRange fetch account data with rage [lower, count)
func AccountsWithRange(lower, count int) ([]AccountModel, error) {
	var result []AccountModel
	var err error
	var rawResult orm.RawSeter

	o := GetOrmObject()
	sqlQuery := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", AccountTable, count, lower)
	rawResult = o.Raw(sqlQuery)

	_, err = rawResult.QueryRows(&result)
	if err != nil {

		beego.Error(err)
		return nil, err
	}

	return result, nil
}

// AllAccounts fetch all account
func AllAccounts() ([]AccountModel, error) {

	return AccountsWithRange(0, -1)
}

// UpdateAccount [WIP] update account's content
func UpdateAccount(id IndexType, params map[string]interface{}) error {

	_, qs := GetQuerySeterWithTable(AccountTable)

	_, err := AccountWithID(id)
	if err != nil {

		beego.Debug("could not find account: %d.", id)

		return err
	}

	beego.Info(params)
	_, err = qs.Filter("id", fmt.Sprintf("%d", id)).Update(params)
	if err != nil {

		beego.Error("[orm] fails to update account")

		return err
	}

	return nil
}

// DeleteAccount [DONE] delete account with id
func DeleteAccount(id IndexType) error {

	o, _ := GetQuerySeterWithTable(AccountTable)

	acc := &AccountModel{ID: id}
	_, err := o.Delete(acc)

	return err
}
