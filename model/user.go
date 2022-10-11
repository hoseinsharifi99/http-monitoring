package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	Urls     []Url  `gorm:"foreignkey:user_id"`
}

func HashPassword(pass string) (string, error) {
	if len(pass) <= 0 {
		return "", errors.New("pass cant be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash), err
}

func (user *User) ValidatePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}
