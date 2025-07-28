package domain

import (
	"time"
)

type Merchant struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	NameMerchant string    `json:"name_merchant"`
	Lat          string    `json:"lat"`
	Long         string    `json:"long"`
	OpenHour     string    `json:"open_hour"`
	CloseHour    string    `json:"close_hour"`
	Status       string    `json:"status"` // e.g. "open", "closed"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// string id = 1;
// string user_id = 2;
// string name_merchant = 3;
// string lat = 4;
// string long = 5;
// string open_hour = 6;
// string close_hour = 7;
// google.protobuf.Timestamp created_at = 8;
// google.protobuf.Timestamp updated_at = 9;
