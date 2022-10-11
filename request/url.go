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
