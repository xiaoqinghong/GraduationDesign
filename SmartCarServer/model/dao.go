package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/xiaoqinghong/SmartCarServer/conf"
)

var db *gorm.DB
var err error

type IDao interface {
	GetTable() string
}

func init() {
	db, err = gorm.Open("mysql", conf.GetConfig().MySQL.GetConnectionStr())
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
}

func DB() *gorm.DB {
	return db
}

func CloseDB() error {
	return db.Close()
}
