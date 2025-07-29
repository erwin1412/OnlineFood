package app_test

import (
	"auth-service/internal/app"
	"auth-service/internal/domain"
	"context"
	"errors"
	"testing"
)

type KafkaProducer struct{}

func (k *KafkaProducer) Publish(message interface{}) error {
	return nil
}

// Producer defines the contract
type Producer interface {
	Publish(message interface{}) error
}

// ✅ Mock Repo
type mockAuthRepo struct {
	GetByEmailFn func(ctx context.Context, email string) (*domain.User, error)
	CreateFn     func(ctx context.Context, user *domain.User) (*domain.User, error)
}

func (m *mockAuthRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return m.GetByEmailFn(ctx, email)
}

func (m *mockAuthRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return m.CreateFn(ctx, user)
}

// ✅ Mock Hasher
type mockHasher struct {
	HashFn  func(password string) (string, error)
	CheckFn func(password, hash string) bool
}

func (m *mockHasher) Hash(password string) (string, error) {
	return m.HashFn(password)
}

func (m *mockHasher) Check(password, hash string) bool {
	return m.CheckFn(password, hash)
}

// ✅ Mock JWT
type mockJWT struct {
	GenerateTokenFn func(userID, email string) (string, error)
}

func (m *mockJWT) GenerateToken(userID, email string) (string, error) {
	return m.GenerateTokenFn(userID, email)
}

func TestRegister_Success(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockAuthRepo{
		GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, nil
		},
		CreateFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			user.ID = "new-id"
			return user, nil
		},
	}

	mockHasher := &mockHasher{
		HashFn: func(password string) (string, error) {
			return "hashed", nil
		},
	}

	jwt := &mockJWT{}
	fakeProducer := &app.NoopProducer{}

	appSvc := app.NewAuthApp(mockRepo, mockHasher, jwt, fakeProducer)

	user, err := appSvc.Register(ctx, "Erwin", "erwin@mail.com", "pass", "admin", "08", "Jakarta", "1.1", "2.2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.ID == "" || user.Name != "Erwin" {
		t.Errorf("unexpected user: %+v", user)
	}
}

func TestRegister_EmailExists(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockAuthRepo{
		GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return &domain.User{ID: "exists"}, nil
		},
	}

	fakeProducer := &app.NoopProducer{}

	appSvc := app.NewAuthApp(mockRepo, nil, nil, fakeProducer)

	_, err := appSvc.Register(ctx, "X", "x@mail.com", "123", "role", "0", "addr", "0", "0")
	if !errors.Is(err, app.ErrEmailExist) {
		t.Errorf("expected ErrEmailExist, got %v", err)
	}
}

func TestLogin_Success(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockAuthRepo{
		GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return &domain.User{ID: "1", Email: email, Password: "hashed"}, nil
		},
	}

	hasher := &mockHasher{
		CheckFn: func(password, hash string) bool {
			return true
		},
	}

	jwt := &mockJWT{
		GenerateTokenFn: func(userID, email string) (string, error) {
			return "token123", nil
		},
	}

	fakeProducer := &app.NoopProducer{}

	appSvc := app.NewAuthApp(mockRepo, hasher, jwt, fakeProducer)

	token, err := appSvc.Login(ctx, "x@mail.com", "pass")
	if err != nil || token != "token123" {
		t.Fatalf("unexpected login result: token=%s err=%v", token, err)
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockAuthRepo{
		GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return &domain.User{ID: "1", Password: "hashed"}, nil
		},
	}

	hasher := &mockHasher{
		CheckFn: func(password, hash string) bool {
			return false
		},
	}

	fakeProducer := &app.NoopProducer{}

	appSvc := app.NewAuthApp(mockRepo, hasher, nil, fakeProducer)

	_, err := appSvc.Login(ctx, "x@mail.com", "wrong")
	if !errors.Is(err, app.ErrUnauthorized) {
		t.Errorf("expected ErrUnauthorized, got %v", err)
	}
}

func TestLogin_NotFound(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockAuthRepo{
		GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, nil
		},
	}

	fakeProducer := &app.NoopProducer{}

	appSvc := app.NewAuthApp(mockRepo, nil, nil, fakeProducer)

	_, err := appSvc.Login(ctx, "x@mail.com", "123")
	if !errors.Is(err, app.ErrUnauthorized) {
		t.Errorf("expected ErrUnauthorized, got %v", err)
	}
}
