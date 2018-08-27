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
	AccountTable string = "account"
)

// IndexType all data type of Index
type IndexType int64

// RuleType about account rule in FFQATracking platform
type RuleType int64

const (
	ruleAdmin RuleType = 0
	ruleUser  RuleType = 1
	ruleGuest RuleType = 2
)

// JobType about job type in FF
type JobType int64

const (
	jobManager      JobType = 0
	jobLeader       JobType = 1
	jobDeveloper    JobType = 2
	jobQATester     JobType = 3
	jobProductOwner JobType = 4
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

// InstallAdminAccount install the admin account, ONLY ONCE
func InstallAdminAccount() {
	acc, err := AddAccount("admin", "root@farfetch.com")
	if err == nil {
		beego.Error(err)
		return
	}

	acc.Rule = ruleAdmin
	acc.Job = jobManager
	acc.Pwd = utils.MD5("admin")

	UpdateAccount(acc.ID, map[string]interface{}{
		"Rule": acc.Rule,
		"Job":  acc.Job,
		"Pwd":  acc.Pwd,
	})
}

// AddAccount insert new account with name and email
func AddAccount(name string, email string) (*AccountModel, error) {

	o := orm.NewOrm()
	account := &AccountModel{Name: name,
		Email:  email,
		Create: time.Now(),
	}

	filterErr := o.QueryTable(AccountTable).Filter("email", email).One(account)
	if filterErr == nil { // account has already existed
		return nil, filterErr
	}

	_, insertErr := o.Insert(account)
	if insertErr != nil {
		return nil, insertErr
	}
	return nil, nil
}

// AccountWithUname get account with uname
func AccountWithUname(uname string) (*AccountModel, error) {

	o := orm.NewOrm()
	acc := &AccountModel{Name: uname}

	filterErr := o.QueryTable(AccountTable).Filter("name", uname).One(acc)
	if filterErr == nil {

		return nil, filterErr
	}

	return acc, nil
}

// AccountWithEmail get account with email
func AccountWithEmail(email string) (*AccountModel, error) {

	o := orm.NewOrm()
	acc := &AccountModel{Email: email}

	filterErr := o.QueryTable(AccountTable).Filter("email", email).One(acc)
	if filterErr == nil {

		return acc, nil
	}

	return nil, filterErr
}

// AccountWithID get account with id
func AccountWithID(id IndexType) (*AccountModel, error) {

	o := orm.NewOrm()
	acc := &AccountModel{ID: id}

	filterErr := o.QueryTable(AccountTable).Filter("id", fmt.Sprintf("%d", id)).One(acc)
	if filterErr == nil {

		return acc, nil
	}

	return nil, filterErr
}

// UpdateAccount [WIP] update account's content
func UpdateAccount(id IndexType, params map[string]interface{}) error {

	o := orm.NewOrm()
	qs := o.QueryTable(AccountTable)

	_, err := AccountWithID(id)
	if err != nil {

		beego.Debug("could not find account: %d.", id)

		return err
	}

	_, err = qs.Filter("id", fmt.Sprintf("%d", id)).Update(params)
	if err != nil {

		beego.Error("[orm] fails to update account")

		return err
	}

	return nil
}

// DeleteAccount [DONE] delete account with id
func DeleteAccount(id IndexType) error {

	o := orm.NewOrm()

	acc := &AccountModel{ID: id}
	_, err := o.Delete(acc)

	return err
}
