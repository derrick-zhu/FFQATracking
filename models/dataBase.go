package models

import (
	"os"
	"path"

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

	if !com.IsExist(_dbFileName) {

		os.MkdirAll(path.Dir(_dbFileName), os.ModePerm)
		os.Create(_dbFileName)
	}
	orm.RegisterDataBase("default", _dbSqlite3Drv, _dbFileName, 10)
	orm.RegisterDriver(_dbSqlite3Drv, orm.DRSqlite)
	orm.RegisterModel(new(AccountModel), new(BugModel))
}
