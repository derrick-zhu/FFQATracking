package models

import (
	"FFQATracking/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.AddFuncMap("PropertyInAccount", PropertyInAccount)
	beego.AddFuncMap("AccountForIDInArray", AccountForIDInArray)
}

// PropertyInAccount ge property value for account
func PropertyInAccount(property string, account *AccountModel) string {

	return utils.FieldInObject(property, account).(string)
}

// AccountForIndex fetch account in account array
func AccountForIndex(accounts *[]AccountModel, index int64) *AccountModel {
	if index < 0 || int(index) >= len(*accounts) {
		return nil
	}
	return &(*accounts)[index]
}

// AccountForIDInArray fetch account object in account array
func AccountForIDInArray(accounts *[]AccountModel, id int64) *AccountModel {
	for _, acc := range *accounts {
		if acc.ID == id {
			return &acc
		}
	}
	return nil
}
