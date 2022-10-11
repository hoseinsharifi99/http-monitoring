package handler

import (
	"ec/auth"
	"ec/db_manager"
	"ec/monitor"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Handler struct {
	dm  *db_manager.DbInstance
	sch *monitor.Scheduler
	ech *echo.Echo
}

func Newhandler(dm *db_manager.DbInstance, sch *monitor.Scheduler) *Handler {
	h := &Handler{dm: dm, sch: sch, ech: echo.New()}
	h.defineRout()
	return h
}

func (h *Handler) defineRout() {
	h.ech.Use(auth.JWT())

	auth.AddToWhiteList("/users/login", "POST")
	auth.AddToWhiteList("/users", "POST")

	//USER
	h.ech.POST("/users", h.SignUp)
	h.ech.POST("/users/login", h.Login)

	//URL
	h.ech.GET("/urls", h.FetchUrls)
	h.ech.POST("/url", h.CreateUrl)
	h.ech.GET("/urls/:urlID", h.GetUrlStats)
}

func (h *Handler) Start() {
	h.ech.Logger.Fatal(h.ech.Start(":8000"))
}

func extractID(c echo.Context) uint {
	e := c.Get("user").(*jwt.Token)
	claims := e.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
