package app

import (
	"context"
	"food-service/internal/domain"
	"time"
)

type FoodApp struct {
	FoodRepo domain.FoodRepository
}

func NewFoodApp(repo domain.FoodRepository) *FoodApp {
	return &FoodApp{
		FoodRepo: repo,
	}
}
func (a *FoodApp) GetById(ctx context.Context, id string) (*domain.Food, error) {
	if id == "" {
		return nil, ErrValidation
	}
	return a.FoodRepo.GetById(ctx, id)
}
func (a *FoodApp) Create(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	if food.Name == "" || food.Price <= 0 {
		return nil, ErrValidation
	}
	food.CreatedAt = time.Now()
	food.UpdatedAt = time.Now()
	return a.FoodRepo.Create(ctx, food)
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
