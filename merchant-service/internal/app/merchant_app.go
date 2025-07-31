package app

import (
	"context"
	"merchant-service/internal/domain"
	"time"
)

type MerchantApp struct {
	MerchantRepo domain.MerchantRepository
}

func NewMerchantApp(repo domain.MerchantRepository) *MerchantApp {
	return &MerchantApp{
		MerchantRepo: repo,
	}
}
func (a *MerchantApp) GetById(ctx context.Context, id string) (*domain.Merchant, error) {
	if id == "" {
		return nil, ErrValidation
	}
	return a.MerchantRepo.GetById(ctx, id)
}
func (a *MerchantApp) Create(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	if merchant.NameMerchant == "" || merchant.Lat == "" || merchant.Long == "" {
		return nil, ErrValidation
	}
	merchant.CreatedAt = time.Now()
	merchant.UpdatedAt = time.Now()
	return a.MerchantRepo.Create(ctx, merchant)
}
func (a *MerchantApp) GetAll(ctx context.Context) ([]*domain.Merchant, error) {
	return a.MerchantRepo.GetAll(ctx)
}
func (a *MerchantApp) Update(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	if merchant.ID == "" || merchant.NameMerchant == "" || merchant.Lat == "" || merchant.Long == "" {
		return nil, ErrValidation
	}
	merchant.UpdatedAt = time.Now()
	return a.MerchantRepo.Update(ctx, merchant)
}
func (a *MerchantApp) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrValidation
	}
	return a.MerchantRepo.Delete(ctx, id)
}

func (a *MerchantApp) GetMerchantByUserId(ctx context.Context, userId string) (*domain.Merchant, error) {
	if userId == "" {
		return nil, ErrValidation
	}
	return a.MerchantRepo.GetMerchantByUserId(ctx, userId)
}
