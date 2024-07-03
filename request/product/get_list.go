package product

import "time"

type (
	GetListRequest struct {
		Keyword  string     `form:"keyword"`
		Name     *string    `form:"name"`
		Page     int        `form:"page"`
		Limit    int        `form:"limit"`
		Sort     string     `form:"sort"`
		Sequence int        `form:"sequence"`
		Date     *time.Time `form:"date"`
	}
	ListResponse struct {
		Name      string    `json:"name"`
		Image     string    `json:"image"`
		Price     int       `json:"price"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Quantity  int       `json:"quantity"`
		Sequence  int       `json:"sequence"`
	}
)
