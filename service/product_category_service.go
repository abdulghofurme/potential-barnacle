package service

import (
	"context"

	"abdulghofur.me/pshamo-go/model/web"
)

type ProductCategoryService interface {
	Create(ctx context.Context, productCategoryRequest web.ProductCategoryCreateRequest) web.ProductCategoryResponse
	Update(ctx context.Context, productCategoryRequest web.ProductCategoryUpdateRequest) web.ProductCategoryResponse
	Delete(ctx context.Context, productCategoryId string) web.ProductCategoryResponse
	FindById(ctx context.Context, productCategoryId string) web.ProductCategoryResponse
	FindAll(ctx context.Context) []web.ProductCategoryResponse
}
