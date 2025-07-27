package domain

import "context"

type TransactionService interface {
	GetById(ctx context.Context, id string, userId string) (*Transaction, error)
	Create(ctx context.Context, transaction *Transaction) (*Transaction, error)
	GetAll(ctx context.Context, userId string) ([]*Transaction, error)
	Update(ctx context.Context, transaction *Transaction) (*Transaction, error)
	Delete(ctx context.Context, id string, userId string) error
}

type TransactionDetailService interface {
	Create(ctx context.Context, detail *TransactionDetail) (*TransactionDetail, error)
	GetByTransactionID(ctx context.Context, transactionID string) ([]*TransactionDetail, error)
	DeleteByTransactionID(ctx context.Context, transactionID string) error
}
