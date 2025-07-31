package domain

import (
	"context"
	"errors"
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
	GetMerchantByUserId(ctx context.Context, userId string) (*Merchant, error)
}

// ErrMerchantNotFound

var (
	ErrMerchantNotFound = errors.New(" merchant tidak ditemukan ")
)
