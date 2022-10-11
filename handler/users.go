package handler

import (
	"ec/model"
	"ec/request"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func bindToAuthRequest(c echo.Context) (*request.UserReq, error) {
	var userAuth = &request.UserReq{}
	if err := c.Bind(userAuth); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "error binding user", err)
	}
	return userAuth, nil
}

func (h *Handler) SignUp(c echo.Context) error {

	req, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		log.Printf("cont validate user data %v", err)
		return echo.ErrBadRequest
	}

	user := &model.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	user.Password, _ = model.HashPassword(user.Password)
	if err := h.dm.AddUser(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not add user to database", err)
	}

	return c.JSON(http.StatusCreated, request.CreateResponseUser(user))
}

func (h *Handler) Login(c echo.Context) error {

	req, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	user := &model.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	u, err := h.dm.GetUserByUserName(user.UserName)
	if err != nil {
		log.Printf("user doesnt exitst %v", err)
		return echo.ErrInternalServerError
	}
	if !u.ValidatePassword(user.Password) {
		log.Printf("password is incorrect %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalide password")
	}
	return c.JSON(http.StatusOK, request.CreateResponseUser(u))
}
