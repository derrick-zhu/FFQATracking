package models

import (
	"os"
	"path"

	"github.com/astaxie/beego"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
)

const (
	_dbFileName   = "data/ffqatracking.db"
	_dbSqlite3Drv = "sqlite3"
)

// IndexType all data type of Index
type IndexType int64

// RegisterDB register and init the DB
func RegisterDB(force bool) {

	bInit := !force && com.IsExist(_dbFileName)
	if !bInit {
		os.MkdirAll(path.Dir(_dbFileName), os.ModePerm)
		os.Create(_dbFileName)
	}

	orm.RegisterDataBase("default", _dbSqlite3Drv, _dbFileName, 10)
	orm.RegisterDriver(_dbSqlite3Drv, orm.DRSqlite)

	// 自动建表
	beego.Info("create db table `default`...")
	orm.RunSyncdb("default", false, true)

	if !bInit {
		beego.Info("setting up and initialize...")
		InstallAdminAccount()
	}
}

// GetOrmObject get orm object
func GetOrmObject() orm.Ormer {

	return orm.NewOrm()
}

// GetQuerySeterWithTable generate a new db QuerySeter object with the name of the db table
func GetQuerySeterWithTable(dbTable string) (orm.Ormer, orm.QuerySeter) {

	o := GetOrmObject()
	return o, o.QueryTable(dbTable)
}
