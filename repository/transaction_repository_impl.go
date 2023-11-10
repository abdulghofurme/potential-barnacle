package repository

import (
	"context"
	"database/sql"
	"errors"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
)

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

type TransactionRepositoryImpl struct{}

func (repository *TransactionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `insert into 
	transactions(id, user_id, address, payment_method, status, total_price, shipping_price) 
	values(?, ?, ?, ?, ?, ?, ?)`

	_, err := tx.ExecContext(
		ctx,
		SQL,
		transaction.Id,
		transaction.UserId,
		transaction.Address,
		transaction.PaymentMethod,
		transaction.Status,
		transaction.TotalPrice,
		transaction.ShippingPrice,
	)
	helper.PanicIfError(err)

	transaction, _ = repository.FindById(ctx, tx, transaction.Id)

	return transaction
}

func (repository *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `update transactions set address=?, payment_method=?, status=?, updated_at=? where id=?`
	_, err := tx.ExecContext(
		ctx,
		SQL,
		transaction.Address,
		transaction.PaymentMethod,
		transaction.Status,
		transaction.UpdatedAt,
		transaction.Id,
	)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) {
	SQL := `update transactions set deleted_at=? where id=?`
	_, err := tx.ExecContext(
		ctx,
		SQL,
		transaction.DeletedAt,
		transaction.Id,
	)
	helper.PanicIfError(err)
}

func (repository *TransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionId string) (domain.Transaction, error) {
	SQL := `select id, address, payment_method, total_price, shipping_price, status, created_at, updated_at, deleted_at where id=?`
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		transactionId,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	transaction := domain.Transaction{}
	if rows.Next() {
		err = rows.Scan(
			&transaction.Id,
			&transaction.Address,
			&transaction.PaymentMethod,
			&transaction.TotalPrice,
			&transaction.ShippingPrice,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
		)
		helper.PanicIfError(err)

		return transaction, nil
	}
	return transaction, errors.New("transaction tidak ditemukan")
}

func (repository *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction {
	SQL := `select id, address, payment_method, total_price, shipping_price, status, created_at, updated_at, deleted_at where deleted_at is null`
	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err = rows.Scan(
			&transaction.Id,
			&transaction.Address,
			&transaction.PaymentMethod,
			&transaction.TotalPrice,
			&transaction.ShippingPrice,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
		)
		helper.PanicIfError(err)

		transactions = append(transactions, transaction)
	}
	return transactions
}
