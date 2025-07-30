package domain

import (
	"time"
)

type Courier struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Lat           string    `json:"lat"`
	Long          string    `json:"long"`
	VehicleNumber string    `json:"vehicle_number"`
	Status        string    `json:"status"` // e.g., "available", "busy", "offline"
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
