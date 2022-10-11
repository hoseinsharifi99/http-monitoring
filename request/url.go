package request

import (
	"ec/model"
	"time"
)

type UrlReq struct {
	Address string `json:"address"`
}

type UrlResponse struct {
	ID           int       `json:"id"`
	URL          string    `json:"url"`
	UserID       uint      `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	Threshold    int       `json:"threshold"`
	FailedTimes  int       `json:"failed_times"`
	SuccessTimes int       `json:"success_times"`
}

type UrlListResponse struct {
	URLs     []*UrlResponse `json:"urls"`
	UrlCount int            `json:"url_count"`
}

func NewUrlListResponse(list []model.Url) *UrlListResponse {
	urls := make([]*UrlResponse, 0)
	for i := range list {
		urls = append(urls, NewUrlResponse(&list[i]))
	}
	return &UrlListResponse{
		URLs:     urls,
		UrlCount: len(urls),
	}
}

func NewUrlResponse(url *model.Url) *UrlResponse {
	return &UrlResponse{
		ID:           int(url.ID),
		URL:          url.Address,
		UserID:       url.UserId,
		Threshold:    url.Threshold,
		FailedTimes:  url.FailedTimes,
		SuccessTimes: url.SuccessTimes,
		CreatedAt:    url.Model.CreatedAt,
	}
}

type requestListResponse struct {
	URL           string             `json:"url"`
	RequestsCount int                `json:"requests_count"`
	Requests      []*requestResponse `json:"requests"`
}

type requestResponse struct {
	ResultCode int       `json:"result_code"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewRequestListResponse(list []model.Request, url string) *requestListResponse {
	resp := new(requestListResponse)
	resp.Requests = make([]*requestResponse, len(list))

	for i := range list {
		resp.Requests[i] = &requestResponse{
			ResultCode: list[i].Result,
			CreatedAt:  list[i].CreatedAt,
		}
	}
	resp.URL = url
	resp.RequestsCount = len(list)
	return resp
}
