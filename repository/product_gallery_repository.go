package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

type ProductGalleryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, productGallery domain.ProductGallery) domain.ProductGallery
	Update(ctx context.Context, tx *sql.Tx, productGallery domain.ProductGallery) domain.ProductGallery
	Delete(ctx context.Context, tx *sql.Tx, productGallery domain.ProductGallery)
	FindById(ctx context.Context, tx *sql.Tx, productGalleryId string) (domain.ProductGallery, error)
	FindByProductId(ctx context.Context, tx *sql.Tx, productId string) []domain.ProductGallery
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductGallery
}
