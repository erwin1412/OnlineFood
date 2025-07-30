package app_test

import (
	"context"
	"errors"
	"food-service/internal/app"
	"food-service/internal/domain"
	"testing"
	"time"
)

// Mock Repo
type mockFoodRepo struct {
	GetByIdFn func(ctx context.Context, id string) (*domain.Food, error)
	CreateFn  func(ctx context.Context, food *domain.Food) (*domain.Food, error)
	GetAllFn  func(ctx context.Context) ([]*domain.Food, error)
	UpdateFn  func(ctx context.Context, food *domain.Food) (*domain.Food, error)
	DeleteFn  func(ctx context.Context, id string) error
}

func (m *mockFoodRepo) GetById(ctx context.Context, id string) (*domain.Food, error) {
	return m.GetByIdFn(ctx, id)
}

func (m *mockFoodRepo) Create(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	return m.CreateFn(ctx, food)
}

func (m *mockFoodRepo) GetAll(ctx context.Context) ([]*domain.Food, error) {
	return m.GetAllFn(ctx)
}

func (m *mockFoodRepo) Update(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	return m.UpdateFn(ctx, food)
}

func (m *mockFoodRepo) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}

// ✅ Mock Producer
type mockProducer struct{}

func (m *mockProducer) Publish(message interface{}) error {
	return nil
}

func TestFoodApp_Create_Success(t *testing.T) {
	mockRepo := &mockFoodRepo{
		CreateFn: func(ctx context.Context, food *domain.Food) (*domain.Food, error) {
			food.ID = "F1"
			return food, nil
		},
	}

	appSvc := app.NewFoodApp(mockRepo, &mockProducer{})

	food := &domain.Food{
		Name:       "Burger",
		Price:      20000,
		MerchantID: "M1",
	}

	result, err := appSvc.Create(context.Background(), food)
	if err != nil || result.ID == "" {
		t.Errorf("unexpected result: %+v, err: %v", result, err)
	}
}

func TestFoodApp_Create_Invalid(t *testing.T) {
	appSvc := app.NewFoodApp(nil, &mockProducer{})

	food := &domain.Food{
		Name:  "",
		Price: 0,
	}

	_, err := appSvc.Create(context.Background(), food)
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestFoodApp_GetById_Success(t *testing.T) {
	mockRepo := &mockFoodRepo{
		GetByIdFn: func(ctx context.Context, id string) (*domain.Food, error) {
			return &domain.Food{ID: id, Name: "Nasi Goreng"}, nil
		},
	}

	appSvc := app.NewFoodApp(mockRepo, &mockProducer{})
	result, err := appSvc.GetById(context.Background(), "F1")
	if err != nil || result.ID != "F1" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestFoodApp_GetById_EmptyId(t *testing.T) {
	appSvc := app.NewFoodApp(nil, &mockProducer{})

	_, err := appSvc.GetById(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestFoodApp_GetAll(t *testing.T) {
	mockRepo := &mockFoodRepo{
		GetAllFn: func(ctx context.Context) ([]*domain.Food, error) {
			return []*domain.Food{{ID: "F1"}, {ID: "F2"}}, nil
		},
	}

	appSvc := app.NewFoodApp(mockRepo, &mockProducer{})
	result, err := appSvc.GetAll(context.Background())
	if err != nil || len(result) != 2 {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestFoodApp_Update_Success(t *testing.T) {
	mockRepo := &mockFoodRepo{
		UpdateFn: func(ctx context.Context, food *domain.Food) (*domain.Food, error) {
			food.UpdatedAt = time.Now()
			return food, nil
		},
	}

	appSvc := app.NewFoodApp(mockRepo, &mockProducer{})
	food := &domain.Food{
		ID:    "F1",
		Name:  "Pizza",
		Price: 50000,
	}

	result, err := appSvc.Update(context.Background(), food)
	if err != nil || result.ID != "F1" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestFoodApp_Update_Invalid(t *testing.T) {
	appSvc := app.NewFoodApp(nil, &mockProducer{})

	food := &domain.Food{
		ID:    "",
		Name:  "",
		Price: 0,
	}

	_, err := appSvc.Update(context.Background(), food)
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestFoodApp_Delete_Success(t *testing.T) {
	mockRepo := &mockFoodRepo{
		DeleteFn: func(ctx context.Context, id string) error {
			return nil
		},
	}

	appSvc := app.NewFoodApp(mockRepo, &mockProducer{})
	err := appSvc.Delete(context.Background(), "F1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFoodApp_Delete_Invalid(t *testing.T) {
	appSvc := app.NewFoodApp(nil, &mockProducer{})
	err := appSvc.Delete(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}
