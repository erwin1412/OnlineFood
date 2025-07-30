package domain

import (
	"time"
)

type Cart struct {
	ID         string    `json:"id"`
	MerchantID string    `json:"merchant_id"`
	FoodID     string    `json:"food_id"`
	UserID     string    `json:"user_id"`
	Qty        int64     `json:"qty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
