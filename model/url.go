package model

import (
	"gorm.io/gorm"
	"net/http"
)

type Url struct {
	gorm.Model
	UserId       uint
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

func (url *Url) SendRequest() (*Request, error) {
	resp, err := http.Get(url.Address)
	request := new(Request)
	request.UrlID = url.ID
	if err != nil {
		return request, err
	}
	request.Result = resp.StatusCode

	return request, nil
}
