package domain

import (
	"context"
)

type CartService interface {
	GetById(ctx context.Context, id string, userId string) (*Cart, error)
	Create(ctx context.Context, cart *Cart) (*Cart, error)
	GetAll(ctx context.Context, userId string) ([]*Cart, error)
	Update(ctx context.Context, cart *Cart) (*Cart, error)
	Delete(ctx context.Context, id string, userId string) error
	DeleteAll(ctx context.Context, userId string) error // ✅ Tambah
}
