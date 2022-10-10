package request

import (
	"ec/auth"
	"ec/model"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func CreateResponseUser(u *model.User) *UserResponse {
	token, _ := auth.GenerateJWTToken(u.ID)
	resUser := &UserResponse{
		UserName: u.UserName,
		Token:    token,
	}

	return resUser
}

func (s UserReq) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.UserName, validation.Required, is.UTFLetterNumeric),
		validation.Field(&s.Password, validation.Required, is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("student validation failled %v", err)
	}
	return nil
}
