package domain

import (
	"time"
)

// ✅ Detail transaksi
type TransactionDetail struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	FoodID        string    `json:"food_id"`
	MerchantID    string    `json:"merchant_id"`
	Qty           int64     `json:"qty"`
	Price         int64     `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ✅ Header transaksi, sekarang punya details
type Transaction struct {
	ID         string               `json:"id"`
	UserID     string               `json:"user_id"`
	CourierID  string               `json:"courier_id"`
	MerchantID string               `json:"merchant_id"`
	Total      int64                `json:"total"`
	Status     string               `json:"status"`
	SnapToken  string               `json:"snap_token"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	Details    []*TransactionDetail `json:"details"` // Embed details!
}
