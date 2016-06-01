package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() error {
	var err error = nil
	db, err = gorm.Open("mysql", Conf.DatabasePath)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	db.SingularTable(true)
	return nil
}
