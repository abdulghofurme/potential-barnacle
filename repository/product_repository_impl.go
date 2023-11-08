package repository

import (
	"context"
	"database/sql"
	"errors"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
)

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

type ProductRepositoryImpl struct{}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := `insert into products(id, name, price, description, tags, category_id) values(?, ?, ?, ?, ?, ?)`
	_, err := tx.ExecContext(ctx, SQL, product.Id, product.Name, product.Price, product.Description, product.Tags, product.CategoryId)
	helper.PanicIfErrof(err)

	product, _ = repository.FindById(ctx, tx, product.Id)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := `update products set name=?, price=?, description=?, tags=?, category_id=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Description, product.Tags, product.CategoryId, product.Id)
	helper.PanicIfErrof(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := `update products set deleted_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, product.DeletedAt.Time, product.Id)
	helper.PanicIfErrof(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	SQL := `select 
	p.id, p.name, p.price, p.description, p.tags, p.created_at, p.updated_at, p.deleted_at,
	pc.id, pc.name
	from products p join product_categories pc on p.category_id=pc.id 
	where p.id=?`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfErrof(err)
	defer rows.Close()

	product := domain.Product{
		Category: domain.ProductCategory{},
	}
	for rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.Tags,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
			&product.Category.Id,
			&product.Category.Name,
		)
		helper.PanicIfErrof(err)

		return product, nil
	}

	return product, errors.New("product tidak ditemukan")

}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := `select 
	p.id, p.name, p.price, p.description, p.tags, p.created_at, p.updated_at, p.deleted_at,
	pc.id, pc.name
	from products p join product_categories pc on p.category_id=pc.id 
	where p.deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErrof(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{
			Category: domain.ProductCategory{},
		}

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.Tags,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
			&product.Category.Id,
			&product.Category.Name,
		)
		helper.PanicIfErrof(err)

		products = append(products, product)
	}

	return products

}
