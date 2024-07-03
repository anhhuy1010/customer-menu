package product

import "time"

type (
	GetDetailUri struct {
		Uuid string `uri:"uuid"`
	}
	GetDetailResponse struct {
		Price       int       `json:"price"`
		Image       string    `json:"image"`
		Name        string    `json:"name"`
		Sequence    int       `json:"sequence"`
		Quantity    int       `json:"quantity"`
		Description string    `json:"description"`
		Gallery     []string  `json:"gallery"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}
)
