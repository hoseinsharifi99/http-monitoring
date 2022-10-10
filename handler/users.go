package handler

import (
	"ec/db_manager"
	"ec/model"
	"ec/request"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	req := new(request.UserReq)

	if err := c.BodyParser(req); err != nil {
		log.Printf("cont load student data %v", err)
		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Printf("cont validate student data %v", err)
		return fiber.ErrBadRequest
	}

	user := &model.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	user.Password, _ = model.HashPassword(user.Password)
	if err := h.dm.AddUser(user); err != nil {
		if errors.Is(err, db_manager.ErrUserDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "User already exist")
		}
		log.Printf("cand save user %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(http.StatusCreated).JSON(request.CreateResponseUser(user))
}
