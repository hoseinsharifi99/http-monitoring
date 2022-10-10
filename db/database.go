package db

import (
	"ec/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDatabase(name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		log.Fatal("failled to Connect to database")
		os.Exit(2)
	}

	log.Println("connected successfully to databsae")
	log.Println("running migration")

	db.AutoMigrate(model.User{}, model.Url{}, model.Request{})

	return db

}
