package app_test

import (
	"context"
	"courier-service/internal/app"
	"courier-service/internal/domain"
	"errors"
	"testing"
	"time"
)

// Mock CourierRepository
type mockCourierRepo struct {
	GetByIdFn       func(ctx context.Context, id string) (*domain.Courier, error)
	CreateFn        func(ctx context.Context, courier *domain.Courier) (*domain.Courier, error)
	GetByLongLatFn  func(ctx context.Context, lat, long string) (*domain.Courier, error)
	UpdateLongLatFn func(ctx context.Context, id, lat, long string) (*domain.Courier, error)
	DeleteFn        func(ctx context.Context, id string) error
	GetAllFn        func(ctx context.Context) ([]*domain.Courier, error)
}

func (m *mockCourierRepo) GetById(ctx context.Context, id string) (*domain.Courier, error) {
	return m.GetByIdFn(ctx, id)
}

func (m *mockCourierRepo) Create(ctx context.Context, courier *domain.Courier) (*domain.Courier, error) {
	return m.CreateFn(ctx, courier)
}

func (m *mockCourierRepo) GetByLongLat(ctx context.Context, lat, long string) (*domain.Courier, error) {
	return m.GetByLongLatFn(ctx, lat, long)
}

func (m *mockCourierRepo) UpdateLongLat(ctx context.Context, id, lat, long string) (*domain.Courier, error) {
	return m.UpdateLongLatFn(ctx, id, lat, long)
}

func (m *mockCourierRepo) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}

func (m *mockCourierRepo) GetAll(ctx context.Context) ([]*domain.Courier, error) {
	return m.GetAllFn(ctx)
}

func TestCourierApp_Create_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := &mockCourierRepo{
		CreateFn: func(ctx context.Context, c *domain.Courier) (*domain.Courier, error) {
			c.ID = "1"
			return c, nil
		},
	}

	appSvc := app.NewCourierApp(mockRepo)

	courier := &domain.Courier{
		UserID:        "U123",
		Lat:           "-6.2",
		Long:          "106.8",
		VehicleNumber: "B1234XYZ",
	}

	created, err := appSvc.Create(ctx, courier)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if created.ID == "" || created.UserID != "U123" {
		t.Errorf("unexpected courier: %+v", created)
	}
}

func TestCourierApp_Create_Invalid(t *testing.T) {
	appSvc := app.NewCourierApp(nil)

	courier := &domain.Courier{
		UserID: "", // invalid
	}

	_, err := appSvc.Create(context.Background(), courier)
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestCourierApp_GetById_Success(t *testing.T) {
	mockRepo := &mockCourierRepo{
		GetByIdFn: func(ctx context.Context, id string) (*domain.Courier, error) {
			return &domain.Courier{ID: id, UserID: "U1"}, nil
		},
	}
	appSvc := app.NewCourierApp(mockRepo)

	result, err := appSvc.GetById(context.Background(), "1")
	if err != nil || result.ID != "1" {
		t.Errorf("unexpected result: %+v, err: %v", result, err)
	}
}

func TestCourierApp_GetById_EmptyId(t *testing.T) {
	appSvc := app.NewCourierApp(nil)

	_, err := appSvc.GetById(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestCourierApp_GetByLongLat_Success(t *testing.T) {
	mockRepo := &mockCourierRepo{
		GetByLongLatFn: func(ctx context.Context, lat, long string) (*domain.Courier, error) {
			return &domain.Courier{Lat: lat, Long: long}, nil
		},
	}

	appSvc := app.NewCourierApp(mockRepo)
	result, err := appSvc.GetByLongLat(context.Background(), "-6.2", "106.8")
	if err != nil || result.Lat != "-6.2" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestCourierApp_GetByLongLat_Empty(t *testing.T) {
	appSvc := app.NewCourierApp(nil)

	_, err := appSvc.GetByLongLat(context.Background(), "", "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestCourierApp_UpdateLongLat_Success(t *testing.T) {
	mockRepo := &mockCourierRepo{
		UpdateLongLatFn: func(ctx context.Context, id, lat, long string) (*domain.Courier, error) {
			return &domain.Courier{ID: id, Lat: lat, Long: long, UpdatedAt: time.Now()}, nil
		},
	}

	appSvc := app.NewCourierApp(mockRepo)
	result, err := appSvc.UpdateLongLat(context.Background(), "1", "-6.3", "107")
	if err != nil || result.ID != "1" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestCourierApp_UpdateLongLat_Invalid(t *testing.T) {
	appSvc := app.NewCourierApp(nil)

	_, err := appSvc.UpdateLongLat(context.Background(), "", "", "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestCourierApp_Delete_Success(t *testing.T) {
	mockRepo := &mockCourierRepo{
		DeleteFn: func(ctx context.Context, id string) error {
			return nil
		},
	}

	appSvc := app.NewCourierApp(mockRepo)
	err := appSvc.Delete(context.Background(), "1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCourierApp_Delete_Invalid(t *testing.T) {
	appSvc := app.NewCourierApp(nil)

	err := appSvc.Delete(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got %v", err)
	}
}

func TestCourierApp_GetAll_Success(t *testing.T) {
	mockRepo := &mockCourierRepo{
		GetAllFn: func(ctx context.Context) ([]*domain.Courier, error) {
			return []*domain.Courier{
				{ID: "1"}, {ID: "2"},
			}, nil
		},
	}

	appSvc := app.NewCourierApp(mockRepo)
	result, err := appSvc.GetAll(context.Background())
	if err != nil || len(result) != 2 {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}
