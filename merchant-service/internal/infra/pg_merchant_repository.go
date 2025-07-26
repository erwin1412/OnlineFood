package infra

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"merchant-service/internal/domain"

	"github.com/google/uuid"
)

// type pgFoodRepository struct {
// 	db *sql.DB
// }

// func NewPgFoodRepository(db *sql.DB) *pgFoodRepository {
// 	if db == nil {
// 		log.Fatal("Postgres DB is nil")
// 	}
// 	return &pgFoodRepository{db}
// }

type pgMerchantRepository struct {
	db *sql.DB
}

func NewPgMerchantRepository(db *sql.DB) *pgMerchantRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &pgMerchantRepository{db}
}

// Implement the MerchantRepository interface methods here
func (r *pgMerchantRepository) GetById(ctx context.Context, id string) (*domain.Merchant, error) {
	var merchant domain.Merchant
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, name_merchant, lat, long, open_hour, close_hour, created_at, updated_at
		FROM merchants
		WHERE id = $1 LIMIT 1
	`, id).Scan(
		&merchant.ID, &merchant.UserID, &merchant.NameMerchant, &merchant.Lat,
		&merchant.Long, &merchant.OpenHour, &merchant.CloseHour, &merchant.CreatedAt,
		&merchant.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No merchant found
		}
		return nil, err
	}
	return &merchant, nil
}
func (r *pgMerchantRepository) Create(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO merchants
		(id, user_id, name_merchant, lat, long, open_hour, close_hour, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`,
		id, merchant.UserID, merchant.NameMerchant, merchant.Lat,
		merchant.Long, merchant.OpenHour, merchant.CloseHour,
		merchant.CreatedAt, merchant.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	merchant.ID = id
	return merchant, nil
}
func (r *pgMerchantRepository) GetAll(ctx context.Context) ([]*domain.Merchant, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, name_merchant, lat, long, open_hour, close_hour, created_at, updated_at
		FROM merchants
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var merchants []*domain.Merchant
	for rows.Next() {
		var merchant domain.Merchant
		if err := rows.Scan(&merchant.ID, &merchant.UserID, &merchant.NameMerchant,
			&merchant.Lat, &merchant.Long, &merchant.OpenHour,
			&merchant.CloseHour, &merchant.CreatedAt, &merchant.UpdatedAt); err != nil {
			return nil, err
		}
		merchants = append(merchants, &merchant)
	}
	return merchants, nil
}
func (r *pgMerchantRepository) Update(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE merchants
		SET user_id = $1, name_merchant = $2, lat = $3, long = $4,
			open_hour = $5, close_hour = $6, updated_at = $7
		WHERE id = $8
	`, merchant.UserID, merchant.NameMerchant, merchant.Lat,
		merchant.Long, merchant.OpenHour, merchant.CloseHour,
		merchant.UpdatedAt, merchant.ID)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}
func (r *pgMerchantRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM merchants
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}
