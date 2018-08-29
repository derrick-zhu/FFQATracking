package models

import (
	"os"
	"path"

	"github.com/astaxie/beego"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_dbFileName   = "data/ffqatracking.db"
	_dbSqlite3Drv = "sqlite3"
)

// RegisterDB register and init the DB
func RegisterDB() {

	bInit := com.IsExist(_dbFileName)
	if !bInit {
		os.MkdirAll(path.Dir(_dbFileName), os.ModePerm)
		os.Create(_dbFileName)
	}

	orm.RegisterDataBase("default", _dbSqlite3Drv, _dbFileName, 10)
	orm.RegisterDriver(_dbSqlite3Drv, orm.DRSqlite)
	//orm.RegisterModel(new(AccountModel), new(BugModel))

	// 自动建表
	beego.Info("create db table `default`...")
	orm.RunSyncdb("default", false, true)

	if !bInit {
		beego.Info("setting up and initialize...")
		InstallAdminAccount()
	}
}

func GetORMWithTable(db_table string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(db_table)
	return o
}
