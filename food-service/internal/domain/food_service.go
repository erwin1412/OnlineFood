package domain

import "context"

type FoodService interface {
	GetById(ctx context.Context, id string) (*Food, error)
	Create(ctx context.Context, food *Food) (*Food, error)
	GetAll(ctx context.Context) ([]*Food, error)
	Update(ctx context.Context, food *Food) (*Food, error)
	Delete(ctx context.Context, id string) error
}
