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
	CourierClient         domain.CourierClient
	MerchantClient        domain.MerchantClient
}

func NewTransactionApp(
	txRepo domain.TransactionRepository,
	detailRepo domain.TransactionDetailRepository,
	cartRepo domain.CartRepository,
	midtransClient *payments.MidtransClient, // Add MidtransClient to constructor
	courierClient domain.CourierClient, // Add CourierClient to constructor
	merchantClient domain.MerchantClient, // Add MerchantClient to constructor
) *TransactionApp {
	return &TransactionApp{
		TransactionRepo:       txRepo,
		TransactionDetailRepo: detailRepo,
		CartRepo:              cartRepo,
		MidtransClient:        midtransClient,
		CourierClient:         courierClient,
		MerchantClient:        merchantClient,
	}
}
func (a *TransactionApp) Create(ctx context.Context, tx *domain.Transaction, details []*domain.TransactionDetail, userID string) (*domain.Transaction, error) {
	if tx.UserID == "" || tx.MerchantID == "" || tx.Total <= 0 || tx.Status == "" {
		return nil, ErrValidation
	}

	// 1️⃣ Get Merchant lat/long
	merchant, err := a.MerchantClient.GetById(ctx, tx.MerchantID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Cari courier terdekat
	nearestCourier, err := a.CourierClient.FindNearest(ctx, merchant.Lat, merchant.Long)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Isi CourierID hasil FindNearest
	tx.CourierID = nearestCourier.ID

	// 4️⃣ ID & timestamp
	tx.ID = uuid.NewString()
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()

	// 5️⃣ Simpan header transaksi
	createdTx, err := a.TransactionRepo.Create(ctx, tx)
	if err != nil {
		return nil, err
	}

	// 6️⃣ Buat detail transaksi
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

	// 7️⃣ Generate Snap token (jika ada)
	if a.MidtransClient != nil {
		snapToken, err := a.MidtransClient.CreateSnapToken(
			createdTx.ID,
			createdTx.Total,
			"Test User",
			"erwin14120824@google.com",
		)
		if err != nil {
			return nil, err
		}
		createdTx.SnapToken = snapToken
	}

	// 8️⃣ Update transaksi dengan Snap token
	createdTx.UpdatedAt = time.Now()
	updatedTx, err := a.TransactionRepo.Update(ctx, createdTx)
	if err != nil {
		return nil, err
	}

	// 9️⃣ Hapus cart
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
