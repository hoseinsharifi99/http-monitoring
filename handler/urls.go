package handler

import (
	"ec/model"
	"ec/request"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func bindToUrlCreateRequest(c echo.Context) (*request.UrlReq, error) {
	request := &request.UrlReq{}
	if err := c.Bind(request); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	return request, nil
}

func (h *Handler) CreateUrl(c echo.Context) error {

	userID := extractID(c)

	req, err := bindToUrlCreateRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, " invalid url", err)
	}

	url := &model.Url{
		UserId:  userID,
		Address: req.Address,
	}

	if err := h.dm.AddUrl(url); err != nil {
		log.Printf("cand save user %v", err)
		return echo.ErrInternalServerError
	}
	/////////////////////////////////
	//            ADD MONITORING
	/////////////////////////////////
	return c.JSON(http.StatusCreated, "URL created successfully")
}
