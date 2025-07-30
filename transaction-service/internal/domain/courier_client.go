// domain/courier_client.go
package domain

import (
	"context"
	"time"
)

type CourierClient interface {
	FindNearest(ctx context.Context, lat, long string) (*Courier, error)
}

type Courier struct {
	ID            string
	UserID        string
	Lat           string
	Long          string
	VehicleNumber string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
