package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction)
	FindById(ctx context.Context, tx *sql.Tx, transactionId string) (domain.Transaction, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction
}
