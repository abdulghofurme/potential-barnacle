package service

import (
	"context"
	"database/sql"
	"time"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/repository"
	"github.com/google/uuid"
)

func NewProductCategoryService(productCategoryRepository repository.ProductCategoryRepository, DB *sql.DB) ProductCategoryService {
	return &ProductCategoryServiceImpl{
		ProductCategoryRepository: productCategoryRepository,
		DB:                        DB,
	}
}

type ProductCategoryServiceImpl struct {
	ProductCategoryRepository repository.ProductCategoryRepository
	DB                        *sql.DB
}

func (service *ProductCategoryServiceImpl) Create(ctx context.Context, productCategoryRequest web.ProductCategoryCreateRequest) web.ProductCategoryResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	existingProductCategories := service.ProductCategoryRepository.FindByName(ctx, tx, productCategoryRequest.Name)
	if len(existingProductCategories) > 0 {
		panic("name sudah digunakan")
	}

	productCategory := domain.ProductCategory{
		Id:   uuid.NewString(),
		Name: productCategoryRequest.Name,
	}

	productCategory = service.ProductCategoryRepository.Create(ctx, tx, productCategory)

	return helper.ToProductCategoryResponse(productCategory)
}

func (service *ProductCategoryServiceImpl) Update(ctx context.Context, productCategoryRequest web.ProductCategoryUpdateRequest) web.ProductCategoryResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	productCategory, err := service.ProductCategoryRepository.FindById(ctx, tx, productCategoryRequest.Id)
	if err != nil {
		panic(err)
	}
	if productCategory.DeletedAt.Valid {
		panic("product category tidak lagi aktif")
	}

	if productCategoryRequest.Name != "" {
		productCategory.Name = productCategoryRequest.Name
	}

	existingProductCategories := service.ProductCategoryRepository.FindByName(ctx, tx, productCategoryRequest.Name)
	if (len(existingProductCategories) == 1 && existingProductCategories[0].Id != productCategory.Id) || len(existingProductCategories) > 1 {
		panic("name sudah digunakan")
	}

	productCategory = service.ProductCategoryRepository.Update(ctx, tx, productCategory)
	return helper.ToProductCategoryResponse(productCategory)
}

func (service *ProductCategoryServiceImpl) Delete(ctx context.Context, productCategoryId string) web.ProductCategoryResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	productCategory, err := service.ProductCategoryRepository.FindById(ctx, tx, productCategoryId)
	if err != nil {
		panic(err)
	}

	if productCategory.DeletedAt.Valid {
		panic("product category tidak lagi aktif")
	}

	productCategory.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	service.ProductCategoryRepository.Delete(ctx, tx, productCategory)
	return helper.ToProductCategoryResponse(productCategory)
}

func (service *ProductCategoryServiceImpl) FindById(ctx context.Context, productCategoryId string) web.ProductCategoryResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	productCategory, err := service.ProductCategoryRepository.FindById(ctx, tx, productCategoryId)
	if err != nil {
		panic(err)
	}
	if productCategory.DeletedAt.Valid {
		panic("product category tidak lagi aktif")
	}

	return helper.ToProductCategoryResponse(productCategory)
}

func (service *ProductCategoryServiceImpl) FindAll(ctx context.Context) []web.ProductCategoryResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	productCategories := service.ProductCategoryRepository.FindAll(ctx, tx)

	var productCategoriesResponse []web.ProductCategoryResponse
	for _, pproductCategory := range productCategories {
		productCategoriesResponse = append(productCategoriesResponse, helper.ToProductCategoryResponse(pproductCategory))
	}

	return productCategoriesResponse
}
