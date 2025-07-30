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

type PgMerchantRepository struct {
	db *sql.DB
}

func NewPgMerchantRepository(db *sql.DB) *PgMerchantRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &PgMerchantRepository{db}
}

// Implement the MerchantRepository interface methods here
func (r *PgMerchantRepository) GetById(ctx context.Context, id string) (*domain.Merchant, error) {
	var merchant domain.Merchant
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, name_merchant,alamat, lat, long, open_hour, close_hour, status, created_at, updated_at
		FROM merchants
		WHERE id = $1 LIMIT 1
	`, id).Scan(
		&merchant.ID, &merchant.UserID, &merchant.NameMerchant, &merchant.Alamat, &merchant.Lat,
		&merchant.Long, &merchant.OpenHour, &merchant.CloseHour, &merchant.Status, &merchant.CreatedAt,
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
func (r *PgMerchantRepository) Create(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO merchants
		(id, user_id, name_merchant, alamat ,lat, long, open_hour, close_hour,status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 , $10 , $11)
	`,
		id, merchant.UserID, merchant.NameMerchant, merchant.Alamat, merchant.Lat,
		merchant.Long, merchant.OpenHour, merchant.CloseHour, merchant.Status,
		merchant.CreatedAt, merchant.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	merchant.ID = id

	// select from user where id = $1 , and update the role = merchant
	// err = r.db.QueryRowContext(ctx, `
	// UPDATE users
	// SET role = 'merchant'
	// WHERE id = $1
	// `, merchant.UserID).Scan()
	// if err != nil {
	// 	log.Println("Error updating user role to merchant:", err)
	// 	return nil, err
	// }

	return merchant, nil
}
func (r *PgMerchantRepository) GetAll(ctx context.Context) ([]*domain.Merchant, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, name_merchant,alamat, lat, long, open_hour, close_hour, status , created_at, updated_at
		FROM merchants
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var merchants []*domain.Merchant
	for rows.Next() {
		var merchant domain.Merchant
		if err := rows.Scan(&merchant.ID, &merchant.UserID, &merchant.NameMerchant, &merchant.Alamat,
			&merchant.Lat, &merchant.Long, &merchant.OpenHour,
			&merchant.CloseHour, &merchant.Status, &merchant.CreatedAt, &merchant.UpdatedAt); err != nil {
			return nil, err
		}
		merchants = append(merchants, &merchant)
	}
	return merchants, nil
}
func (r *PgMerchantRepository) Update(ctx context.Context, merchant *domain.Merchant) (*domain.Merchant, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE merchants
		SET user_id = $1, name_merchant = $2, alamat = $3,lat = $4, long = $5,
			open_hour = $6, close_hour = $7, updated_at = $8 , status = $9
		WHERE id = $10
	`, merchant.UserID, merchant.NameMerchant, merchant.Alamat, merchant.Lat,
		merchant.Long, merchant.OpenHour, merchant.CloseHour,
		merchant.UpdatedAt, merchant.Status, merchant.ID)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}
func (r *PgMerchantRepository) Delete(ctx context.Context, id string) error {

	if id == "" {
		return errors.New("merchant ID cannot be empty")
	}
	// set the user role to user
	var userID string

	// set merchant.UserID = userID
	err := r.db.QueryRowContext(ctx, `
		SELECT user_id FROM merchants WHERE id = $1
	`, id).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("merchant not found")
		}
	}

	_, err = r.db.ExecContext(ctx, `
		DELETE FROM merchants
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	// set the user role to user
	_, err = r.db.ExecContext(ctx, `
		UPDATE users
		SET role = 'user'
		WHERE id = $1
	`, userID)
	if err != nil {
		log.Println("Error updating user role to user:", err)
	}

	return nil
}
