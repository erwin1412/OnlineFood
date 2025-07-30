package app_test

import (
	"context"
	"errors"
	"merchant-service/internal/app"
	"merchant-service/internal/domain"
	"testing"
	"time"
)

// Mock MerchantRepo
type mockMerchantRepo struct {
	GetByIdFn func(ctx context.Context, id string) (*domain.Merchant, error)
	CreateFn  func(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error)
	GetAllFn  func(ctx context.Context) ([]*domain.Merchant, error)
	UpdateFn  func(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error)
	DeleteFn  func(ctx context.Context, id string) error
}

func (m *mockMerchantRepo) GetById(ctx context.Context, id string) (*domain.Merchant, error) {
	return m.GetByIdFn(ctx, id)
}

func (m *mockMerchantRepo) Create(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	return m.CreateFn(ctx, merchant)
}

func (m *mockMerchantRepo) GetAll(ctx context.Context) ([]*domain.Merchant, error) {
	return m.GetAllFn(ctx)
}

func (m *mockMerchantRepo) Update(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	return m.UpdateFn(ctx, merchant)
}

func (m *mockMerchantRepo) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}

func TestMerchantApp_Create_Success(t *testing.T) {
	mockRepo := &mockMerchantRepo{
		CreateFn: func(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
			merchant.ID = "M1"
			return merchant, nil
		},
	}

	appSvc := app.NewMerchantApp(mockRepo)

	merchant := &domain.Merchant{
		NameMerchant: "Toko A",
		Lat:          "1.2345",
		Long:         "6.7890",
	}

	result, err := appSvc.Create(context.Background(), merchant)
	if err != nil || result.ID == "" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestMerchantApp_Create_Invalid(t *testing.T) {
	appSvc := app.NewMerchantApp(nil)

	merchant := &domain.Merchant{
		NameMerchant: "",
		Lat:          "",
		Long:         "",
	}

	_, err := appSvc.Create(context.Background(), merchant)
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestMerchantApp_GetById_Success(t *testing.T) {
	mockRepo := &mockMerchantRepo{
		GetByIdFn: func(ctx context.Context, id string) (*domain.Merchant, error) {
			return &domain.Merchant{ID: id, NameMerchant: "Toko B"}, nil
		},
	}

	appSvc := app.NewMerchantApp(mockRepo)

	result, err := appSvc.GetById(context.Background(), "M1")
	if err != nil || result.ID != "M1" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestMerchantApp_GetById_Invalid(t *testing.T) {
	appSvc := app.NewMerchantApp(nil)

	_, err := appSvc.GetById(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestMerchantApp_GetAll(t *testing.T) {
	mockRepo := &mockMerchantRepo{
		GetAllFn: func(ctx context.Context) ([]*domain.Merchant, error) {
			return []*domain.Merchant{{ID: "M1"}, {ID: "M2"}}, nil
		},
	}

	appSvc := app.NewMerchantApp(mockRepo)

	result, err := appSvc.GetAll(context.Background())
	if err != nil || len(result) != 2 {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestMerchantApp_Update_Success(t *testing.T) {
	mockRepo := &mockMerchantRepo{
		UpdateFn: func(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
			merchant.UpdatedAt = time.Now()
			return merchant, nil
		},
	}

	appSvc := app.NewMerchantApp(mockRepo)

	merchant := &domain.Merchant{
		ID:           "M1",
		NameMerchant: "Toko Updated",
		Lat:          "1.1111",
		Long:         "2.2222",
	}

	result, err := appSvc.Update(context.Background(), merchant)
	if err != nil || result.ID != "M1" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestMerchantApp_Update_Invalid(t *testing.T) {
	appSvc := app.NewMerchantApp(nil)

	merchant := &domain.Merchant{
		ID:           "",
		NameMerchant: "",
		Lat:          "",
		Long:         "",
	}

	_, err := appSvc.Update(context.Background(), merchant)
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestMerchantApp_Delete_Success(t *testing.T) {
	mockRepo := &mockMerchantRepo{
		DeleteFn: func(ctx context.Context, id string) error {
			return nil
		},
	}

	appSvc := app.NewMerchantApp(mockRepo)

	err := appSvc.Delete(context.Background(), "M1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestMerchantApp_Delete_Invalid(t *testing.T) {
	appSvc := app.NewMerchantApp(nil)

	err := appSvc.Delete(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}
