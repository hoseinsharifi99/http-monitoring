package handler

import (
	"ec/model"
	"ec/request"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
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
	h.sch.Mnt.AddURL([]*model.Url{url})
	return c.JSON(http.StatusCreated, "URL created successfully")
}

func (h *Handler) FetchUrls(c echo.Context) error {
	userID := extractID(c)

	urls, err := h.dm.GetUrlsByUserId(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
		log.Println(err)
	}

	resp := request.NewUrlListResponse(urls)
	return c.JSON(http.StatusOK, resp)

}

func (h *Handler) GetUrlStats(c echo.Context) error {
	userID := extractID(c)
	urlID, err := strconv.Atoi(c.Param("urlID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalide urlid", err)
	}

	var url *model.Url
	url, err = h.dm.GetUrlById(uint(urlID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "can not load url by id", err)
	}

	if url.UserId != userID {
		return echo.NewHTTPError(http.StatusBadRequest, "you not alow ", err)
	}

	return c.JSON(http.StatusOK, request.NewRequestListResponse(url.Requests, url.Address))
}
