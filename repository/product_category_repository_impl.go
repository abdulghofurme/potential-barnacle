package repository

import (
	"context"
	"database/sql"
	"errors"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
)

func NewProductCategoryRepository() ProductCategoryRepository {
	return &ProductCategoryRepositoryImpl{}
}

type ProductCategoryRepositoryImpl struct{}

func (repository *ProductCategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, productCategory domain.ProductCategory) domain.ProductCategory {
	SQL := `insert into product_categories(id, name) values(id, name)`
	_, err := tx.ExecContext(ctx, SQL, productCategory.Id, productCategory.Name)
	helper.PanicIfErrof(err)

	productCategory, _ = repository.FindById(ctx, tx, productCategory.Id)

	return productCategory
}

func (repository *ProductCategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, productCategory domain.ProductCategory) domain.ProductCategory {
	SQL := `update product_categories set name=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, productCategory.Name, productCategory.Id)
	helper.PanicIfErrof(err)

	return productCategory
}

func (repository *ProductCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productCategory domain.ProductCategory) {
	SQL := `update product_categories set deleted_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, productCategory.DeletedAt.Time, productCategory.Id)
	helper.PanicIfErrof(err)
}

func (repository *ProductCategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productCategoryId string) (domain.ProductCategory, error) {
	SQL := `select id, name, created_at, updated_at, deleted_at from product_categories where id=?`
	rows, err := tx.QueryContext(ctx, SQL, productCategoryId)
	helper.PanicIfErrof(err)
	defer rows.Close()

	productCategory := domain.ProductCategory{}
	if rows.Next() {
		err := rows.Scan(
			&productCategory.Id,
			&productCategory.Name,
			&productCategory.CreatedAt,
			&productCategory.UpdatedAt,
			&productCategory.DeletedAt,
		)
		helper.PanicIfErrof(err)
		return productCategory, nil
	}
	return productCategory, errors.New("product category tidak ditemukan")
}

func (repository *ProductCategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductCategory {
	SQL := `select id, name, created_at, updated_at, deleted_at from product_categories where deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErrof(err)
	defer rows.Close()

	var productCategories []domain.ProductCategory
	for rows.Next() {
		productCategory := domain.ProductCategory{}
		err := rows.Scan(
			&productCategory.Id,
			&productCategory.Name,
			&productCategory.CreatedAt,
			&productCategory.UpdatedAt,
			&productCategory.DeletedAt,
		)
		helper.PanicIfErrof(err)
		productCategories = append(productCategories, productCategory)
	}

	return productCategories
}

func (repository *ProductCategoryRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) []domain.ProductCategory {
	SQL := `select id, name from product_categories where name=?`
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfErrof(err)
	defer rows.Close()

	var productCategories []domain.ProductCategory
	for rows.Next() {
		productCategory := domain.ProductCategory{}
		err := rows.Scan(
			&productCategory.Id,
			&productCategory.Name,
		)
		helper.PanicIfErrof(err)
		productCategories = append(productCategories, productCategory)
	}

	return productCategories
}
