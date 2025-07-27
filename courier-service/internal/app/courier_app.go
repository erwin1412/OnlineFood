package app

import (
	"context"
	"courier-service/internal/domain"
	"time"
)

// type MerchantApp struct {
// 	MerchantRepo domain.MerchantRepository
// }

// func NewMerchantApp(repo domain.MerchantRepository) *MerchantApp {
// 	return &MerchantApp{
// 		MerchantRepo: repo,
// 	}
// }

type CourierApp struct {
	CourierRepo domain.CourierRepository
}

func NewCourierApp(repo domain.CourierRepository) *CourierApp {
	return &CourierApp{
		CourierRepo: repo,
	}
}

func (a *CourierApp) GetById(ctx context.Context, id string) (*domain.Courier, error) {
	if id == "" {
		return nil, ErrValidation
	}
	return a.CourierRepo.GetById(ctx, id)
}
func (a *CourierApp) Create(ctx context.Context, courier *domain.Courier) (*domain.Courier, error) {
	if courier.UserID == "" || courier.Lat == "" || courier.Long == "" || courier.VehicleNumber == "" {
		return nil, ErrValidation
	}
	courier.CreatedAt = time.Now()
	courier.UpdatedAt = time.Now()
	return a.CourierRepo.Create(ctx, courier)
}
func (a *CourierApp) GetByLongLat(ctx context.Context, lat, long string) (*domain.Courier, error) {
	if lat == "" || long == "" {
		return nil, ErrValidation
	}
	return a.CourierRepo.GetByLongLat(ctx, lat, long)
}
func (a *CourierApp) UpdateLongLat(ctx context.Context, id, lat, long string) (*domain.Courier, error) {
	if id == "" || lat == "" || long == "" {
		return nil, ErrValidation
	}
	return a.CourierRepo.UpdateLongLat(ctx, id, lat, long)
}
func (a *CourierApp) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrValidation
	}
	return a.CourierRepo.Delete(ctx, id)
}
func (a *CourierApp) GetAll(ctx context.Context) ([]*domain.Courier, error) {
	return a.CourierRepo.GetAll(ctx)
}
