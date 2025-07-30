package domain

import "context"

type CourierService interface {
	GetById(ctx context.Context, id string) (*Courier, error)
	Create(ctx context.Context, courier *Courier) (*Courier, error)
	GetByLongLat(ctx context.Context, lat, long string) (*Courier, error)
	UpdateLongLat(ctx context.Context, id, lat, long string) (*Courier, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*Courier, error)
	FindNearest(ctx context.Context, lat, long string) (*Courier, error)
}
