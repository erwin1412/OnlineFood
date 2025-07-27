package app

import (
	"context"
	"time"
	"transaction-service/internal/domain"
	"transaction-service/pkg/payments"

	"github.com/google/uuid"
)

type TransactionApp struct {
	TransactionRepo       domain.TransactionRepository
	TransactionDetailRepo domain.TransactionDetailRepository
	CartRepo              domain.CartRepository
	MidtransClient        *payments.MidtransClient // Add MidtransClient
}

func NewTransactionApp(
	txRepo domain.TransactionRepository,
	detailRepo domain.TransactionDetailRepository,
	cartRepo domain.CartRepository,
	midtransClient *payments.MidtransClient, // Add MidtransClient to constructor
) *TransactionApp {
	return &TransactionApp{
		TransactionRepo:       txRepo,
		TransactionDetailRepo: detailRepo,
		CartRepo:              cartRepo,
		MidtransClient:        midtransClient,
	}
}

// func (a *TransactionApp) Create(ctx context.Context, tx *domain.Transaction, details []*domain.TransactionDetail, userID string) (*domain.Transaction, error) {
// 	if tx.UserID == "" || tx.CourierID == "" || tx.MerchantID == "" || tx.Total <= 0 || tx.Status == "" {
// 		return nil, ErrValidation
// 	}

// 	tx.ID = uuid.NewString()
// 	tx.CreatedAt = time.Now()
// 	tx.UpdatedAt = time.Now()

// 	// 1️⃣ Buat header transaksi
// 	createdTx, err := a.TransactionRepo.Create(ctx, tx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 2️⃣ Buat detail transaksi
// 	for _, detail := range details {
// 		detail.ID = uuid.NewString()
// 		detail.TransactionID = createdTx.ID
// 		detail.MerchantID = createdTx.MerchantID // ✅ isi dari header
// 		detail.CreatedAt = time.Now()
// 		detail.UpdatedAt = time.Now()

// 		_, err := a.TransactionDetailRepo.Create(ctx, detail)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// 3️⃣ Generate Snap token
// 	// Using UserID as placeholder for customerName and customerEmail
// 	// Replace with actual user details if available via a UserService
// 	snapToken, err := a.MidtransClient.CreateSnapToken(
// 		createdTx.ID,
// 		createdTx.Total,
// 		createdTx.UserID, // Placeholder for customer name
// 		createdTx.UserID, // Placeholder for customer email
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 4️⃣ Update transaction with Snap token
// 	createdTx.SnapToken = snapToken
// 	createdTx.UpdatedAt = time.Now()
// 	updatedTx, err := a.TransactionRepo.Update(ctx, createdTx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 5️⃣ Hapus semua cart user
// 	if err := a.CartRepo.DeleteAll(ctx, userID); err != nil {
// 		return nil, err
// 	}

// 	return updatedTx, nil
// }

func (a *TransactionApp) Create(ctx context.Context, tx *domain.Transaction, details []*domain.TransactionDetail, userID string) (*domain.Transaction, error) {
	if tx.UserID == "" || tx.CourierID == "" || tx.MerchantID == "" || tx.Total <= 0 || tx.Status == "" {
		return nil, ErrValidation
	}

	tx.ID = uuid.NewString()
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()

	// 1️⃣ Buat header transaksi
	createdTx, err := a.TransactionRepo.Create(ctx, tx)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Buat detail transaksi
	for _, detail := range details {
		detail.ID = uuid.NewString()
		detail.TransactionID = createdTx.ID
		detail.CreatedAt = time.Now()
		detail.UpdatedAt = time.Now()

		_, err := a.TransactionDetailRepo.Create(ctx, detail)
		if err != nil {
			return nil, err
		}
	}

	// 3️⃣ Generate Snap token
	// Use hardcoded test values for sandbox testing
	snapToken, err := a.MidtransClient.CreateSnapToken(
		createdTx.ID,
		createdTx.Total,
		"Test User",                // Hardcoded customer name for sandbox
		"erwin14120824@google.com", // Hardcoded valid email for sandbox
	)
	if err != nil {
		return nil, err
	}

	// 4️⃣ Update transaction with Snap token
	createdTx.SnapToken = snapToken
	createdTx.UpdatedAt = time.Now()
	updatedTx, err := a.TransactionRepo.Update(ctx, createdTx)
	if err != nil {
		return nil, err
	}

	// 5️⃣ Hapus semua cart user
	if err := a.CartRepo.DeleteAll(ctx, userID); err != nil {
		return nil, err
	}

	return updatedTx, nil
}

func (a *TransactionApp) GetAll(ctx context.Context, userId string) ([]*domain.Transaction, error) {
	if userId == "" {
		return nil, ErrValidation
	}

	txs, err := a.TransactionRepo.GetAll(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, tx := range txs {
		details, err := a.TransactionDetailRepo.GetByTransactionID(ctx, tx.ID)
		if err != nil {
			return nil, err
		}
		tx.Details = details
	}

	return txs, nil
}

func (a *TransactionApp) GetById(ctx context.Context, id, userId string) (*domain.Transaction, error) {
	if id == "" || userId == "" {
		return nil, ErrValidation
	}

	// Ambil header
	tx, err := a.TransactionRepo.GetById(ctx, id, userId)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, nil
	}

	// Ambil detail & tempelkan
	details, err := a.TransactionDetailRepo.GetByTransactionID(ctx, tx.ID)
	if err != nil {
		return nil, err
	}
	tx.Details = details

	return tx, nil
}

func (a *TransactionApp) Update(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
	if tx.ID == "" || tx.UserID == "" || tx.CourierID == "" || tx.MerchantID == "" || tx.Total <= 0 || tx.Status == "" {
		return nil, ErrValidation
	}

	tx.UpdatedAt = time.Now()

	return a.TransactionRepo.Update(ctx, tx)
}

func (a *TransactionApp) Delete(ctx context.Context, id, userId string) error {
	if id == "" || userId == "" {
		return ErrValidation
	}

	// 1️⃣ Hapus detail dulu
	if err := a.TransactionDetailRepo.DeleteByTransactionID(ctx, id); err != nil {
		return err
	}

	// 2️⃣ Baru hapus header
	return a.TransactionRepo.Delete(ctx, id, userId)
}
