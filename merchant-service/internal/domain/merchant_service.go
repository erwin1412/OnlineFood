package domain

import "context"

type MerchantService interface {
	GetById(ctx context.Context, id string) (*Merchant, error)
	Create(ctx context.Context, merchant *Merchant) (*Merchant, error)
	GetAll(ctx context.Context) ([]*Merchant, error)
	Update(ctx context.Context, merchant *Merchant) (*Merchant, error)
	Delete(ctx context.Context, id string) error
}
