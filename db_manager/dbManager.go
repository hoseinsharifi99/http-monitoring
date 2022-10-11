package db_manager

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database = DbInstance{}

func NewDb(database *gorm.DB) *DbInstance {
	log.Println("JADIDE")
	return &DbInstance{Db: database}
}

var (
	ErrStudentNotFinding = errors.New("there is no student with given id")
	ErrUserDuplicate     = errors.New("student with that id already exist")
)
