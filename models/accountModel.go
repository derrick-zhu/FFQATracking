package models

import (
	"FFQATracking/src/utils"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
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

	UpdateAccount(acc)
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

	return nil, nil
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
func UpdateAccount(account *AccountModel) (*AccountModel, error) {

	o := orm.NewOrm()
	qs := o.QueryTable(AccountTable)

	acc, err := AccountWithID(account.ID)
	if err != nil {
		beego.Debug("could not find account: %s.", account.Email)
		return nil, err
	}

	source := utils.Struct2Map(*account)
	param := orm.Params{}

	for key, value := range source {
		if key == "ID" || key == "Email" || key == "Create" {
			continue
		}

		param[key] = value
	}

	_, err = qs.Filter("id", fmt.Sprintf("%d", account.ID)).Update(param)
	if err != nil {
		beego.Error("[orm] fails to update account")
		return nil, err
	}

	return acc, nil
}

// DeleteAccount [DONE] delete account with id
func DeleteAccount(id IndexType) error {

	o := orm.NewOrm()

	acc := &AccountModel{ID: id}
	_, err := o.Delete(acc)

	return err
}
