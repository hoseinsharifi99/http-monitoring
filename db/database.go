package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database = DbInstance{}

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("myproject.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failled to Connect to database")
		os.Exit(2)
	}

	log.Println("connected successfully to databsae")
	log.Println("running migration")

	db.AutoMigrate()
	log.Println("JADIDE")
	Database = DbInstance{
		Db: db,
	}

}
