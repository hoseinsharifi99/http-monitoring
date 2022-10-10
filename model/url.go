package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	UserId       uint   `gorm:"unique"`
	Address      string `gorm:"unique"`
	Threshold    int
	FailedTimes  int
	SuccessTimes int
	Requests     []Request `gorm:"foreignkey:url_id"`
}

type Request struct {
	gorm.Model
	UrlID  uint
	Result int
}
