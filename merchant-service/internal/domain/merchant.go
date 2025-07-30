package domain

import (
	"time"
)

type Merchant struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	NameMerchant string    `json:"name_merchant"`
	Alamat       string    `json:"alamat"` // Address for geocoding
	Lat          string    `json:"lat"`
	Long         string    `json:"long"`
	OpenHour     string    `json:"open_hour"`
	CloseHour    string    `json:"close_hour"`
	Status       string    `json:"status"` // e.g. "open", "closed"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
