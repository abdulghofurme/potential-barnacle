package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
