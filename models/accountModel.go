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

// IndexType all data type of Index
type IndexType int64

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

	o := orm.NewOrm()
	account := &AccountModel{Name: name,
		Email:  email,
		Create: time.Now(),
	}

	filterErr := o.QueryTable(AccountTable).Filter("email", email).One(account)
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

	o := orm.NewOrm()
	acc := &AccountModel{Name: uname}

	filterErr := o.QueryTable(AccountTable).Filter("name", uname).One(acc)
	if filterErr != nil {
		return nil, filterErr
	}

	return acc, nil
}

// AccountWithEmail get account with email
func AccountWithEmail(email string) (*AccountModel, error) {

	o := orm.NewOrm()
	acc := &AccountModel{Email: email}

	filterErr := o.QueryTable(AccountTable).Filter("email", email).One(acc)
	if filterErr != nil {
		return nil, filterErr
	}

	return acc, nil
}

// AccountWithID get account with id
func AccountWithID(id IndexType) (*AccountModel, error) {

	o := orm.NewOrm()
	acc := &AccountModel{ID: id}

	filterErr := o.QueryTable(AccountTable).Filter("id", fmt.Sprintf("%d", id)).One(acc)
	if filterErr != nil {
		return nil, filterErr
	}

	return acc, nil
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

	o := orm.NewOrm()

	acc := &AccountModel{ID: id}
	_, err := o.Delete(acc)

	return err
}
