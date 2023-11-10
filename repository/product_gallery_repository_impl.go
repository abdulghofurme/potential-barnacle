package repository

import (
	"context"
	"database/sql"
	"errors"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
)

func NewProductGalleryRepository() ProductGalleryRepository {
	return &ProductGalleryRepositoryImpl{}
}

type ProductGalleryRepositoryImpl struct{}

func (repository *ProductGalleryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, productGallery domain.ProductGallery) domain.ProductGallery {
	SQL := `insert into product_galleries(id, url, product_id) values(?, ?, ?)`
	_, err := tx.ExecContext(ctx, SQL, productGallery.Id, productGallery.Url, productGallery.ProductId)
	helper.PanicIfError(err)

	productGallery, _ = repository.FindById(ctx, tx, productGallery.Id)
	return productGallery
}

func (repository *ProductGalleryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, productGallery domain.ProductGallery) domain.ProductGallery {
	SQL := `update product_galleries set url=?, product_id=?, updated_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, productGallery.Url, productGallery.ProductId, productGallery.UpdatedAt, productGallery.Id)
	helper.PanicIfError(err)

	return productGallery
}

func (repository *ProductGalleryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productGallery domain.ProductGallery) {
	SQL := `update product_gallerues set deleted_at=? where id=?`
	_, err := tx.ExecContext(ctx, SQL, productGallery.DeletedAt, productGallery.Id)
	helper.PanicIfError(err)
}

func (repository *ProductGalleryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productGalleryId string) (domain.ProductGallery, error) {
	SQL := `select id, url, product_id, created_at, updated_at, deleted_at from product_galleries where id=?`
	rows, err := tx.QueryContext(ctx, SQL, productGalleryId)
	helper.PanicIfError(err)
	defer rows.Close()

	productGallery := domain.ProductGallery{}
	if rows.Next() {

		err := rows.Scan(
			&productGallery.Id,
			&productGallery.Url,
			&productGallery.ProductId,
			&productGallery.CreatedAt,
			&productGallery.UpdatedAt,
			&productGallery.DeletedAt,
		)
		helper.PanicIfError(err)
		return productGallery, nil
	}
	return productGallery, errors.New("product gallery tidak ditemukan")
}

func (repository *ProductGalleryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductGallery {
	SQL := `select id, url, product_id, created_at, updated_at, deleted_at from product_galleries where deleted_at is null`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var productGalleries []domain.ProductGallery
	for rows.Next() {
		productGallery := domain.ProductGallery{}

		err := rows.Scan(
			&productGallery.Id,
			&productGallery.Url,
			&productGallery.ProductId,
			&productGallery.CreatedAt,
			&productGallery.UpdatedAt,
			&productGallery.DeletedAt,
		)
		helper.PanicIfError(err)
		productGalleries = append(productGalleries, productGallery)
	}

	return productGalleries
}

func (repository *ProductGalleryRepositoryImpl) FindByProductId(ctx context.Context, tx *sql.Tx, productId string) []domain.ProductGallery {
	SQL := `select id, url, product_id, created_at, updated_at, deleted_at from product_galleries where product_id=?`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	var productGalleries []domain.ProductGallery
	for rows.Next() {
		productGallery := domain.ProductGallery{}

		err := rows.Scan(
			&productGallery.Id,
			&productGallery.Url,
			&productGallery.ProductId,
			&productGallery.CreatedAt,
			&productGallery.UpdatedAt,
			&productGallery.DeletedAt,
		)
		helper.PanicIfError(err)
		productGalleries = append(productGalleries, productGallery)
	}

	return productGalleries
}
