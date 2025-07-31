package app

import (
	"context"
	"food-service/internal/domain"
	"time"
)

// Gunakan interface, ini contract
type Producer interface {
	Publish(message interface{}) error
}

type FoodApp struct {
	FoodRepo       domain.FoodRepository
	Producer       Producer              // ✅ interface, tidak terikat struct
	MerchantClient domain.MerchantClient // Add MerchantClient
}

func NewFoodApp(repo domain.FoodRepository, producer Producer, merchantClient domain.MerchantClient) *FoodApp {
	return &FoodApp{
		FoodRepo:       repo,
		Producer:       producer,
		MerchantClient: merchantClient,
	}
}

func (a *FoodApp) GetById(ctx context.Context, id string) (*domain.Food, error) {
	if id == "" {
		return nil, ErrValidation
	}
	return a.FoodRepo.GetById(ctx, id)
}

func (a *FoodApp) Create(ctx context.Context, food *domain.Food, userID string) (*domain.Food, error) {
	if food.Name == "" || food.Price <= 0 {
		return nil, ErrValidation
	}

	merchant, err := a.MerchantClient.GetMerchantByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	if merchant == nil {
		return nil, domain.ErrMerchantNotFound
	}

	food.CreatedAt = time.Now()
	food.UpdatedAt = time.Now()

	createdFood, err := a.FoodRepo.Create(ctx, food)
	if err != nil {
		return nil, err
	}

	if a.Producer != nil {
		_ = a.Producer.Publish(map[string]interface{}{
			"event":       "FoodCreated",
			"id":          createdFood.ID,
			"merchant_id": createdFood.MerchantID,
			"name":        createdFood.Name,
			"price":       createdFood.Price,
			"created_at":  createdFood.CreatedAt,
		})
	}

	return createdFood, nil
}

func (a *FoodApp) GetAll(ctx context.Context) ([]*domain.Food, error) {
	return a.FoodRepo.GetAll(ctx)
}

func (a *FoodApp) Update(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	if food.ID == "" || food.Name == "" || food.Price <= 0 {
		return nil, ErrValidation
	}
	food.UpdatedAt = time.Now()
	return a.FoodRepo.Update(ctx, food)
}

func (a *FoodApp) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrValidation
	}
	return a.FoodRepo.Delete(ctx, id)
}
