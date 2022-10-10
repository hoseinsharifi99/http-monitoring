package handler

import (
	"ec/db_manager"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Handler struct {
	dm  *db_manager.DbInstance
	fib *fiber.App
}

func Newhandler(dm *db_manager.DbInstance) *Handler {
	h := &Handler{dm: dm, fib: fiber.New()}
	h.defineRout()
	return h
}

func (h *Handler) defineRout() {
	h.fib.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello")
	})

	h.fib.Post("/user", h.SignUp)

}

func (h *Handler) Start() {
	if err := h.fib.Listen(":8000"); err != nil {
		log.Println("cant connect to sv")
	}
}
