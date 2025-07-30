package app

import (
	"context"
	"time"
	"transaction-service/internal/domain"

	"github.com/google/uuid"
)

type CartApp struct {
	CartRepo domain.CartRepository
}

func NewCartApp(repo domain.CartRepository) *CartApp {
	return &CartApp{
		CartRepo: repo,
	}
}

func (a *CartApp) Create(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	if cart.MerchantID == "" || cart.FoodID == "" || cart.UserID == "" || cart.Qty <= 0 {
		return nil, ErrValidation
	}

	cart.ID = uuid.NewString()
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()

	return a.CartRepo.Create(ctx, cart)
}

func (a *CartApp) GetAll(ctx context.Context, userId string) ([]*domain.Cart, error) {
	if userId == "" {
		return nil, ErrValidation
	}
	return a.CartRepo.GetAll(ctx, userId)
}

func (a *CartApp) GetById(ctx context.Context, id, userId string) (*domain.Cart, error) {
	if id == "" || userId == "" {
		return nil, ErrValidation
	}
	return a.CartRepo.GetById(ctx, id, userId)
}

func (a *CartApp) Update(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	if cart.ID == "" || cart.MerchantID == "" || cart.FoodID == "" || cart.UserID == "" || cart.Qty <= 0 {
		return nil, ErrValidation
	}

	cart.UpdatedAt = time.Now()

	return a.CartRepo.Update(ctx, cart)
}

func (a *CartApp) Delete(ctx context.Context, id, userId string) error {
	if id == "" || userId == "" {
		return ErrValidation
	}
	return a.CartRepo.Delete(ctx, id, userId)
}

func (a *CartApp) DeleteAll(ctx context.Context, userId string) error {
	if userId == "" {
		return ErrValidation
	}
	return a.CartRepo.DeleteAll(ctx, userId)
}
