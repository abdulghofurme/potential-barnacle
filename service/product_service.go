package service

import (
	"context"

	"abdulghofur.me/pshamo-go/model/web"
)

type ProductService interface {
	Create(ctx context.Context, productRequest web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, productRequest web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productId string) web.ProductResponse
	FindById(ctx context.Context, productId string) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
