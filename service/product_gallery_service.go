package service

import (
	"context"

	"abdulghofur.me/pshamo-go/model/web"
)

type ProductGalleryService interface {
	Create(ctx context.Context, productGalleryRequest web.ProductGalleryCreateRequest) web.ProductGalleryResponse
	Update(ctx context.Context, productGalleryRequest web.ProductGalleryUpdateRequest) web.ProductGalleryResponse
	Delete(ctx context.Context, productGalleryId string) web.ProductGalleryResponse
	FindById(ctx context.Context, productGalleryId string) web.ProductGalleryResponse
	FindAll(ctx context.Context) []web.ProductGalleryResponse
}
