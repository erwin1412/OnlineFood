package infra

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"transaction-service/internal/domain"

	"github.com/google/uuid"
)

type pgTransactionRepository struct {
	db *sql.DB
}
type pgTransactionDetailRepository struct {
	*pgTransactionRepository
}

func NewPgTransactionRepository(db *sql.DB) *pgTransactionRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &pgTransactionRepository{db}
}

// ✅ Create Transaction Header
func (r *pgTransactionRepository) Create(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO transactions
		(id, user_id, courier_id, merchant_id, total, status, snap_token, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`,
		id, tx.UserID, tx.CourierID, tx.MerchantID, tx.Total, tx.Status, tx.SnapToken, tx.CreatedAt, tx.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	tx.ID = id
	return tx, nil
}

// ✅ GetById Transaction Header
func (r *pgTransactionRepository) GetById(ctx context.Context, id string, userId string) (*domain.Transaction, error) {
	var tx domain.Transaction
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, courier_id, merchant_id, total, status, snap_token, created_at, updated_at
		FROM transactions
		WHERE id = $1 AND user_id = $2 LIMIT 1
	`, id, userId).Scan(
		&tx.ID, &tx.UserID, &tx.CourierID, &tx.MerchantID, &tx.Total, &tx.Status, &tx.SnapToken, &tx.CreatedAt, &tx.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &tx, nil
}

// ✅ GetAll Transactions for user
func (r *pgTransactionRepository) GetAll(ctx context.Context, userId string) ([]*domain.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, courier_id, merchant_id, total, status, snap_token, created_at, updated_at
		FROM transactions
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txs []*domain.Transaction

	for rows.Next() {
		var tx domain.Transaction
		if err := rows.Scan(
			&tx.ID, &tx.UserID, &tx.CourierID, &tx.MerchantID, &tx.Total, &tx.Status, &tx.SnapToken, &tx.CreatedAt, &tx.UpdatedAt,
		); err != nil {
			return nil, err
		}
		txs = append(txs, &tx)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return txs, nil
}

// ✅ Update Transaction
func (r *pgTransactionRepository) Update(ctx context.Context, tx *domain.Transaction) (*domain.Transaction, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE transactions
		SET courier_id = $1, merchant_id = $2, total = $3, status = $4, snap_token = $5, updated_at = $6
		WHERE id = $7 AND user_id = $8
	`,
		tx.CourierID, tx.MerchantID, tx.Total, tx.Status, tx.SnapToken, tx.UpdatedAt, tx.ID, tx.UserID,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ✅ Delete Transaction
func (r *pgTransactionRepository) Delete(ctx context.Context, id, userId string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM transactions
		WHERE id = $1 AND user_id = $2
	`, id, userId)
	return err
}

func (r *pgTransactionRepository) CreateDetail(ctx context.Context, detail *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO transaction_details
		(id, transaction_id, food_id, qty, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		id, detail.TransactionID, detail.FoodID, detail.Qty, detail.Price, detail.CreatedAt, detail.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	detail.ID = id
	return detail, nil
}

func (r *pgTransactionRepository) GetByTransactionID(ctx context.Context, transactionID string) ([]*domain.TransactionDetail, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, transaction_id, food_id, qty, price, created_at, updated_at
		FROM transaction_details
		WHERE transaction_id = $1
	`, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []*domain.TransactionDetail

	for rows.Next() {
		var detail domain.TransactionDetail
		if err := rows.Scan(
			&detail.ID, &detail.TransactionID, &detail.FoodID, &detail.Qty, &detail.Price, &detail.CreatedAt, &detail.UpdatedAt,
		); err != nil {
			return nil, err
		}
		details = append(details, &detail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return details, nil
}

func (r *pgTransactionRepository) DeleteByTransactionID(ctx context.Context, transactionID string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM transaction_details
		WHERE transaction_id = $1
	`, transactionID)
	return err
}

func NewPgTransactionDetailRepository(base *pgTransactionRepository) *pgTransactionDetailRepository {
	return &pgTransactionDetailRepository{base}
}

func (r *pgTransactionDetailRepository) Create(ctx context.Context, detail *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	return r.pgTransactionRepository.CreateDetail(ctx, detail)
}

func (r *pgTransactionDetailRepository) GetByTransactionID(ctx context.Context, transactionID string) ([]*domain.TransactionDetail, error) {
	return r.pgTransactionRepository.GetByTransactionID(ctx, transactionID)
}

func (r *pgTransactionDetailRepository) DeleteByTransactionID(ctx context.Context, transactionID string) error {
	return r.pgTransactionRepository.DeleteByTransactionID(ctx, transactionID)
}
