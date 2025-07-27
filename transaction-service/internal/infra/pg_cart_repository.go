package infra

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"transaction-service/internal/domain"

	"github.com/google/uuid"
)

type pgCartRepository struct {
	db *sql.DB
}

func NewPgCartRepository(db *sql.DB) *pgCartRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &pgCartRepository{db}
}

func (r *pgCartRepository) GetById(ctx context.Context, id, userId string) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.db.QueryRowContext(ctx, `
		SELECT id, merchant_id, food_id, user_id, qty, created_at, updated_at
		FROM carts
		WHERE id = $1 AND user_id = $2 LIMIT 1
	`, id, userId).Scan(
		&cart.ID, &cart.MerchantID, &cart.FoodID, &cart.UserID, &cart.Qty, &cart.CreatedAt, &cart.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (r *pgCartRepository) Create(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO carts
		(id, merchant_id, food_id, user_id, qty, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		id, cart.MerchantID, cart.FoodID, cart.UserID, cart.Qty, cart.CreatedAt, cart.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	cart.ID = id
	return cart, nil
}

func (r *pgCartRepository) GetAll(ctx context.Context, userId string) ([]*domain.Cart, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, merchant_id, food_id, user_id, qty, created_at, updated_at
		FROM carts
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []*domain.Cart

	for rows.Next() {
		var cart domain.Cart
		if err := rows.Scan(
			&cart.ID, &cart.MerchantID, &cart.FoodID, &cart.UserID, &cart.Qty, &cart.CreatedAt, &cart.UpdatedAt,
		); err != nil {
			return nil, err
		}
		carts = append(carts, &cart)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *pgCartRepository) Update(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE carts
		SET merchant_id = $1, food_id = $2, qty = $3, updated_at = $4
		WHERE id = $5 AND user_id = $6
	`,
		cart.MerchantID, cart.FoodID, cart.Qty, cart.UpdatedAt, cart.ID, cart.UserID,
	)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *pgCartRepository) Delete(ctx context.Context, id string, userId string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM carts
		WHERE id = $1 AND user_id = $2
	`, id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgCartRepository) DeleteAll(ctx context.Context, userId string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM carts
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return err
	}
	return nil
}
