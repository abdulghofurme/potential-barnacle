package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

type ProductCategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, productCategory domain.ProductCategory) domain.ProductCategory
	Update(ctx context.Context, tx *sql.Tx, productCategory domain.ProductCategory) domain.ProductCategory
	Delete(ctx context.Context, tx *sql.Tx, productCategory domain.ProductCategory)
	FindById(ctx context.Context, tx *sql.Tx, productCategoryId string) (domain.ProductCategory, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductCategory
	FindByName(ctx context.Context, tx *sql.Tx, name string) []domain.ProductCategory
}
