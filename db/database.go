package db

import (
	"ec/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetUpDb(dbName string) *gorm.DB {
	db := NewDb(dbName)
	migrate(db)
	db.LogMode(true)
	return db
}

func NewDb(name string) *gorm.DB {

	db, err := gorm.Open("sqlite3", "./"+name)
	if err != nil {
		fmt.Println("Error in creating database file : ", err)
		return nil
	}
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.User{}, model.Url{}, model.Request{})
}
