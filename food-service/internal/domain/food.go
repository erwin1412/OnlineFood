package domain

import (
	"time"
)

type Food struct {
	ID           string    `json:"id"`
	MerchantID   string    `json:"merchant_id"`
	Name         string    `json:"name_foods"`
	Price        int64     `json:"price"`
	Availability string    `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
