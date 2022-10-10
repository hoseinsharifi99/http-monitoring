package db

import (
	"ec/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database = DbInstance{}

func ConnectDatabase(name string) *DbInstance {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		log.Fatal("failled to Connect to database")
		os.Exit(2)
	}

	log.Println("connected successfully to databsae")
	log.Println("running migration")

	db.AutoMigrate(model.User{}, model.Url{}, model.Request{})

	log.Println("JADIDE")
	return &DbInstance{Db: db}

}
