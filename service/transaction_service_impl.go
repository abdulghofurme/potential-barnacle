package service

import (
	"context"
	"database/sql"
	"time"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewTransactionService(
	transactionRepository repository.TransactionRepository,
	DB *sql.DB,
	validate *validator.Validate,
) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func (service *TransactionServiceImpl) Create(ctx context.Context, transactionRequest web.TransactionCreateRequest) web.TransactionResponse {
	err := service.Validate.Struct(transactionRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction := domain.Transaction{
		Id:            uuid.NewString(),
		UserId:        transactionRequest.UserId,
		Address:       transactionRequest.Address,
		PaymentMethod: transactionRequest.PaymentMethod,
		Status:        transactionRequest.Status,
		TotalPrice:    transactionRequest.TotalPrice,
		ShippingPrice: transactionRequest.ShippingPrice,
	}

	transaction = service.TransactionRepository.Create(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Update(ctx context.Context, transactionRequest web.TransactionUpdateRequest) web.TransactionResponse {
	err := service.Validate.Struct(transactionRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionRequest.Id)
	helper.PanicIfError(err)
	if transaction.DeletedAt.Valid {
		panic("transaction tidak lagi ada")
	}

	transaction.Address = transactionRequest.Address
	transaction.Status = transactionRequest.Status
	transaction.PaymentMethod = transaction.PaymentMethod
	transaction.UpdatedAt = time.Now()

	transaction = service.TransactionRepository.Update(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Delete(ctx context.Context, transactionId string) web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	helper.PanicIfError(err)
	if transaction.DeletedAt.Valid {
		panic("transaction tidak lagi ada")
	}

	transaction.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	service.TransactionRepository.Delete(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindById(ctx context.Context, transactionId string) web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	helper.PanicIfError(err)
	if transaction.DeletedAt.Valid {
		panic("transaction tidak lagi ada")
	}

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindAll(ctx context.Context) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactions := service.TransactionRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	var transactionsResponse []web.TransactionResponse
	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, helper.ToTransactionResponse(transaction))
	}

	return transactionsResponse
}
