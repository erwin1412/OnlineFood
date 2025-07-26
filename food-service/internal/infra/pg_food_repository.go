package infra

import (
	"context"
	"database/sql"
	"errors"
	"food-service/internal/domain"
	"log"

	"github.com/google/uuid"
)

type pgFoodRepository struct {
	db *sql.DB
}

func NewPgFoodRepository(db *sql.DB) *pgFoodRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &pgFoodRepository{db}
}

func (r *pgFoodRepository) GetById(ctx context.Context, id string) (*domain.Food, error) {
	var food domain.Food
	err := r.db.QueryRowContext(ctx, `
		SELECT id, merchant_id, name_foods, price, availability, created_at, updated_at 
		FROM foods
		WHERE id = $1 LIMIT 1
	`, id).Scan(
		&food.ID, &food.MerchantID, &food.Name, &food.Price, &food.Availability, &food.CreatedAt, &food.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No food found
		}
		return nil, err
	}
	return &food, nil
}
func (r *pgFoodRepository) Create(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO foods
		(id, merchant_id, name_foods, price, availability, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		id, food.MerchantID, food.Name, food.Price, food.Availability, food.CreatedAt, food.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	food.ID = id
	return food, nil
}

func (r *pgFoodRepository) GetAll(ctx context.Context) ([]*domain.Food, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, merchant_id, name_foods, price, availability, created_at, updated_at 
		FROM foods
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []*domain.Food
	for rows.Next() {
		var food domain.Food
		if err := rows.Scan(&food.ID, &food.MerchantID, &food.Name, &food.Price, &food.Availability, &food.CreatedAt, &food.UpdatedAt); err != nil {
			return nil, err
		}
		foods = append(foods, &food)
	}
	return foods, nil
}
func (r *pgFoodRepository) Update(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE foods
		SET merchant_id = $1, name_foods = $2, price = $3, availability = $4, updated_at = $5
		WHERE id = $6
	`, food.MerchantID, food.Name, food.Price, food.Availability, food.UpdatedAt, food.ID)
	if err != nil {
		return nil, err
	}
	return food, nil
}
func (r *pgFoodRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM foods
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}
