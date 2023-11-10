package service

import (
	"context"

	"abdulghofur.me/pshamo-go/model/web"
)

type TransactionService interface {
	Create(ctx context.Context, transactionRequest web.TransactionCreateRequest) web.TransactionResponse
	Update(ctx context.Context, transactionRequest web.TransactionUpdateRequest) web.TransactionResponse
	Delete(ctx context.Context, transactionId string) web.TransactionResponse
	FindById(ctx context.Context, transactionId string) web.TransactionResponse
	FindAll(ctx context.Context) []web.TransactionResponse
}
