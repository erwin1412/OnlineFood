package infra

import (
	"context"
	"courier-service/internal/domain"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
)

// type pgMerchantRepository struct {
// 	db *sql.DB
// }

// func NewPgMerchantRepository(db *sql.DB) *pgMerchantRepository {
// 	if db == nil {
// 		log.Fatal("Postgres DB is nil")
// 	}
// 	return &pgMerchantRepository{db}
// }

type pgCourierRepository struct {
	db *sql.DB
}

func NewPgCourierRepository(db *sql.DB) *pgCourierRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &pgCourierRepository{db}
}

// Implement the CourierRepository interface methods here
func (r *pgCourierRepository) GetById(ctx context.Context, id string) (*domain.Courier, error) {
	var courier domain.Courier
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, lat, long, vehicle_number, status, created_at, updated_at
		FROM couriers
		WHERE id = $1 LIMIT 1
	`, id).Scan(
		&courier.ID, &courier.UserID, &courier.Lat, &courier.Long,
		&courier.VehicleNumber, &courier.Status, &courier.CreatedAt,
		&courier.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No courier found
		}
		return nil, err
	}
	return &courier, nil
}
func (r *pgCourierRepository) Create(ctx context.Context, courier *domain.Courier) (*domain.Courier, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO couriers
		(id, user_id, lat, long, vehicle_number, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`,
		id, courier.UserID, courier.Lat, courier.Long,
		courier.VehicleNumber, courier.Status,
		courier.CreatedAt, courier.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	courier.ID = id
	return courier, nil
}
func (r *pgCourierRepository) GetByLongLat(ctx context.Context, lat, long string) (*domain.Courier, error) {
	var courier domain.Courier
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, lat, long, vehicle_number, status, created_at, updated_at
		FROM couriers
		WHERE lat = $1 AND long = $2 LIMIT 1
	`, lat, long).Scan(
		&courier.ID, &courier.UserID, &courier.Lat, &courier.Long,
		&courier.VehicleNumber, &courier.Status, &courier.CreatedAt,
		&courier.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No courier found
		}
		return nil, err
	}
	return &courier, nil
}
func (r *pgCourierRepository) UpdateLongLat(ctx context.Context, id, lat, long string) (*domain.Courier, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE couriers
		SET lat = $1, long = $2, updated_at = NOW()
		WHERE id = $3
	`, lat, long, id)
	if err != nil {
		return nil, err
	}
	// Fetch the updated courier to return
	courier, err := r.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if courier == nil {
		return nil, errors.New("courier not found after update")
	}
	return courier, nil
}
func (r *pgCourierRepository) Delete(ctx context.Context, id string) error {

	if id == "" {
		return errors.New("courier ID cannot be empty")
	}
	var userID string
	// set courier.UserID = userID
	err := r.db.QueryRowContext(ctx, `
		SELECT user_id FROM couriers WHERE id = $1
	`, id).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("courier not found")
		}
	}

	_, err = r.db.ExecContext(ctx, `
		DELETE FROM couriers
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
		return err
	}

	return nil
}
func (r *pgCourierRepository) GetAll(ctx context.Context) ([]*domain.Courier, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, lat, long, vehicle_number, status, created_at, updated_at
		FROM couriers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var couriers []*domain.Courier
	for rows.Next() {
		var courier domain.Courier
		if err := rows.Scan(&courier.ID, &courier.UserID, &courier.Lat,
			&courier.Long, &courier.VehicleNumber, &courier.Status,
			&courier.CreatedAt, &courier.UpdatedAt); err != nil {
			return nil, err
		}
		couriers = append(couriers, &courier)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return couriers, nil
}
