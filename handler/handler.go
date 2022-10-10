package handler

import (
	"ec/db"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Handler struct {
	dm  *db.DbInstance
	fib *fiber.App
}

func Newhandler(dm *db.DbInstance) *Handler {
	h := &Handler{dm: dm, fib: fiber.New()}
	h.defineRout()
	return h
}

func (h *Handler) defineRout() {
	h.fib.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello")
	})
}

func (h *Handler) Start() {
	if err := h.fib.Listen(":8000"); err != nil {
		log.Println("cant connect to sv")
	}
}
