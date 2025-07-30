package domain

import (
	"context"
	"time"
)

type Merchant struct {
	ID           string
	UserID       string
	NameMerchant string
	Alamat       string
	Lat          string
	Long         string
	OpenHour     string
	CloseHour    string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MerchantClient interface {
	GetById(ctx context.Context, id string) (*Merchant, error)
}
