package app_test

import (
	"context"
	"errors"
	"testing"
	"transaction-service/internal/app"
	"transaction-service/internal/domain"
)

// Mock TransactionRepo
type mockTransactionRepo struct {
	CreateFn  func(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error)
	UpdateFn  func(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error)
	GetAllFn  func(ctx context.Context, userId string) ([]*domain.Transaction, error)
	GetByIdFn func(ctx context.Context, id, userId string) (*domain.Transaction, error)
	DeleteFn  func(ctx context.Context, id, userId string) error
}

func (m *mockTransactionRepo) Create(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
	return m.CreateFn(ctx, tx)
}
func (m *mockTransactionRepo) Update(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
	return m.UpdateFn(ctx, tx)
}
func (m *mockTransactionRepo) GetAll(ctx context.Context, userId string) ([]*domain.Transaction, error) {
	return m.GetAllFn(ctx, userId)
}
func (m *mockTransactionRepo) GetById(ctx context.Context, id, userId string) (*domain.Transaction, error) {
	return m.GetByIdFn(ctx, id, userId)
}
func (m *mockTransactionRepo) Delete(ctx context.Context, id, userId string) error {
	return m.DeleteFn(ctx, id, userId)
}

// Mock TransactionDetailRepo
type mockTransactionDetailRepo struct {
	CreateFn                func(ctx context.Context, detail *domain.TransactionDetail) (*domain.TransactionDetail, error)
	GetByTransactionIDFn    func(ctx context.Context, txID string) ([]*domain.TransactionDetail, error)
	DeleteByTransactionIDFn func(ctx context.Context, txID string) error
}

func (m *mockTransactionDetailRepo) Create(ctx context.Context, detail *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	return m.CreateFn(ctx, detail)
}
func (m *mockTransactionDetailRepo) GetByTransactionID(ctx context.Context, txID string) ([]*domain.TransactionDetail, error) {
	return m.GetByTransactionIDFn(ctx, txID)
}
func (m *mockTransactionDetailRepo) DeleteByTransactionID(ctx context.Context, txID string) error {
	return m.DeleteByTransactionIDFn(ctx, txID)
}

// Mock CartRepo
type mockCartRepo struct {
	DeleteAllFn func(ctx context.Context, userId string) error
}

// Create implements domain.CartRepository.
func (m *mockCartRepo) Create(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	panic("unimplemented")
}

// Delete implements domain.CartRepository.
func (m *mockCartRepo) Delete(ctx context.Context, id string, userId string) error {
	panic("unimplemented")
}

// GetAll implements domain.CartRepository.
func (m *mockCartRepo) GetAll(ctx context.Context, userId string) ([]*domain.Cart, error) {
	panic("unimplemented")
}

// GetById implements domain.CartRepository.
func (m *mockCartRepo) GetById(ctx context.Context, id string, userId string) (*domain.Cart, error) {
	panic("unimplemented")
}

// Update implements domain.CartRepository.
func (m *mockCartRepo) Update(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	panic("unimplemented")
}

func (m *mockCartRepo) DeleteAll(ctx context.Context, userId string) error {
	return m.DeleteAllFn(ctx, userId)
}

// Mock MidtransClient
type mockMidtransClient struct {
	CreateSnapTokenFn func(orderID string, total int, customerName, customerEmail string) (string, error)
}

func (m *mockMidtransClient) CreateSnapToken(orderID string, total int, customerName, customerEmail string) (string, error) {
	return m.CreateSnapTokenFn(orderID, total, customerName, customerEmail)
}

func TestTransactionApp_Create_Success(t *testing.T) {
	mockTxRepo := &mockTransactionRepo{
		CreateFn: func(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
			return tx, nil
		},
		UpdateFn: func(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
			return tx, nil
		},
	}
	mockDetailRepo := &mockTransactionDetailRepo{
		CreateFn: func(ctx context.Context, detail *domain.TransactionDetail) (*domain.TransactionDetail, error) {
			return detail, nil
		},
	}
	mockCartRepo := &mockCartRepo{
		DeleteAllFn: func(ctx context.Context, userId string) error {
			return nil
		},
	}

	appSvc := app.NewTransactionApp(mockTxRepo, mockDetailRepo, mockCartRepo, nil)

	tx := &domain.Transaction{
		UserID:     "U1",
		CourierID:  "C1",
		MerchantID: "M1",
		Total:      100000,
		Status:     "pending",
	}

	details := []*domain.TransactionDetail{
		{FoodID: "F1", Qty: 1},
	}

	result, err := appSvc.Create(context.Background(), tx, details, "U1")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Karena MidtransClient nil, SnapToken harus kosong
	if result.SnapToken != "" {
		t.Errorf("expected SnapToken to be empty, got: %s", result.SnapToken)
	}
}

func TestTransactionApp_Create_Invalid(t *testing.T) {
	appSvc := app.NewTransactionApp(nil, nil, nil, nil)

	tx := &domain.Transaction{
		UserID: "",
	}

	_, err := appSvc.Create(context.Background(), tx, nil, "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestTransactionApp_GetAll_Success(t *testing.T) {
	mockTxRepo := &mockTransactionRepo{
		GetAllFn: func(ctx context.Context, userId string) ([]*domain.Transaction, error) {
			return []*domain.Transaction{{ID: "T1"}}, nil
		},
	}
	mockDetailRepo := &mockTransactionDetailRepo{
		GetByTransactionIDFn: func(ctx context.Context, txID string) ([]*domain.TransactionDetail, error) {
			return []*domain.TransactionDetail{{ID: "D1"}}, nil
		},
	}

	appSvc := app.NewTransactionApp(mockTxRepo, mockDetailRepo, nil, nil)

	result, err := appSvc.GetAll(context.Background(), "U1")
	if err != nil || len(result) != 1 || len(result[0].Details) != 1 {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestTransactionApp_GetAll_Invalid(t *testing.T) {
	appSvc := app.NewTransactionApp(nil, nil, nil, nil)

	_, err := appSvc.GetAll(context.Background(), "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestTransactionApp_GetById_Success(t *testing.T) {
	mockTxRepo := &mockTransactionRepo{
		GetByIdFn: func(ctx context.Context, id, userId string) (*domain.Transaction, error) {
			return &domain.Transaction{ID: id}, nil
		},
	}
	mockDetailRepo := &mockTransactionDetailRepo{
		GetByTransactionIDFn: func(ctx context.Context, txID string) ([]*domain.TransactionDetail, error) {
			return []*domain.TransactionDetail{{ID: "D1"}}, nil
		},
	}

	appSvc := app.NewTransactionApp(mockTxRepo, mockDetailRepo, nil, nil)

	result, err := appSvc.GetById(context.Background(), "T1", "U1")
	if err != nil || result.ID != "T1" || len(result.Details) != 1 {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestTransactionApp_GetById_Invalid(t *testing.T) {
	appSvc := app.NewTransactionApp(nil, nil, nil, nil)

	_, err := appSvc.GetById(context.Background(), "", "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestTransactionApp_Update_Success(t *testing.T) {
	mockTxRepo := &mockTransactionRepo{
		UpdateFn: func(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
			return tx, nil
		},
	}

	appSvc := app.NewTransactionApp(mockTxRepo, nil, nil, nil)

	tx := &domain.Transaction{
		ID:         "T1",
		UserID:     "U1",
		CourierID:  "C1",
		MerchantID: "M1",
		Total:      100000,
		Status:     "paid",
	}

	result, err := appSvc.Update(context.Background(), tx)
	if err != nil || result.ID != "T1" {
		t.Errorf("unexpected: %+v, err: %v", result, err)
	}
}

func TestTransactionApp_Update_Invalid(t *testing.T) {
	appSvc := app.NewTransactionApp(nil, nil, nil, nil)

	tx := &domain.Transaction{
		ID: "",
	}

	_, err := appSvc.Update(context.Background(), tx)
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}

func TestTransactionApp_Delete_Success(t *testing.T) {
	mockTxRepo := &mockTransactionRepo{
		DeleteFn: func(ctx context.Context, id, userId string) error {
			return nil
		},
	}
	mockDetailRepo := &mockTransactionDetailRepo{
		DeleteByTransactionIDFn: func(ctx context.Context, txID string) error {
			return nil
		},
	}

	appSvc := app.NewTransactionApp(mockTxRepo, mockDetailRepo, nil, nil)

	err := appSvc.Delete(context.Background(), "T1", "U1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTransactionApp_Delete_Invalid(t *testing.T) {
	appSvc := app.NewTransactionApp(nil, nil, nil, nil)

	err := appSvc.Delete(context.Background(), "", "")
	if !errors.Is(err, app.ErrValidation) {
		t.Errorf("expected ErrValidation, got: %v", err)
	}
}
