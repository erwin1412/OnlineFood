package infra

import (
	"auth-service/internal/domain"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
)

type pgAuthRepository struct {
	db *sql.DB
}

func NewPgAuthRepository(db *sql.DB) *pgAuthRepository {
	if db == nil {
		log.Fatal("Postgres DB is nil")
	}
	return &pgAuthRepository{db}
}

func (r *pgAuthRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {

	var user domain.User

	err := r.db.QueryRowContext(ctx, `SELECT id,name,email,password,role,phone,address,latitude,longitude , created_at, updated_at FROM users WHERE email = $1 LIMIT 1`, email).Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.Phone, &user.Address, &user.Latitude, &user.Longitude, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *pgAuthRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	id := uuid.NewString()
	if id == "" {
		return nil, errors.New("failed to generate UUID")
	}

	_, err := r.db.ExecContext(ctx, `INSERT INTO users(id,name,email,password,role,phone,address,lat,long)`)

	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}
