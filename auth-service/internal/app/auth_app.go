package app

import (
	"auth-service/internal/domain"
	"context"
	"time"
)

type AuthApp struct {
	UserRepo       domain.AuthRepository
	PasswordHasher PasswordHasher
	JWTManager     JWTManager
	Producer       Producer // ✅ GANTI JADI INTERFACE!
}

func NewAuthApp(
	repo domain.AuthRepository,
	hasher PasswordHasher,
	jwt JWTManager,
	producer Producer,
) *AuthApp {
	return &AuthApp{
		UserRepo:       repo,
		PasswordHasher: hasher,
		JWTManager:     jwt,
		Producer:       producer,
	}
}

func (a *AuthApp) Register(
	ctx context.Context,
	name, email, password, role, phone, alamat, latitude, longitude string,
) (*domain.User, error) {

	if name == "" || email == "" || password == "" {
		return nil, ErrValidation
	}

	existing, _ := a.UserRepo.GetByEmail(ctx, email)
	if existing != nil && existing.ID != "" {
		return nil, ErrEmailExist
	}

	hashed, err := a.PasswordHasher.Hash(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:      name,
		Email:     email,
		Password:  hashed,
		Role:      role,
		Phone:     phone,
		Alamat:    alamat,
		Latitude:  latitude,
		Longitude: longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := a.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	if a.Producer != nil {
		_ = a.Producer.Publish(map[string]interface{}{
			"event":      "UserRegistered",
			"id":         createdUser.ID,
			"email":      createdUser.Email,
			"name":       createdUser.Name,
			"created_at": createdUser.CreatedAt,
		})
	}

	return createdUser, nil
}

func (a *AuthApp) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.UserRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return "", ErrUnauthorized
	}

	if !a.PasswordHasher.Check(password, user.Password) {
		return "", ErrUnauthorized
	}

	token, err := a.JWTManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
