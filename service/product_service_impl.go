package service

import (
	"context"
	"database/sql"
	"time"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func (service *ProductServiceImpl) Create(ctx context.Context, productRequest web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(productRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Id:    uuid.NewString(),
		Name:  productRequest.Name,
		Price: productRequest.Price,
		Description: sql.NullString{
			String: productRequest.Description,
			Valid:  true,
		},
		Tags: sql.NullString{
			String: productRequest.Tags,
			Valid:  true,
		},
		CategoryId: productRequest.CategoryId,
	}

	product = service.ProductRepository.Create(
		ctx,
		tx,
		product,
	)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, productRequest web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(productRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productRequest.Id)
	helper.PanicIfError(err)
	if product.DeletedAt.Valid {
		panic("product tidak lagi aktif")
	}

	product.Name = productRequest.Name
	product.Price = productRequest.Price
	product.Description = sql.NullString{
		String: productRequest.Description,
		Valid:  true,
	}
	product.Tags = sql.NullString{
		String: productRequest.Tags,
		Valid:  true,
	}
	product.CategoryId = productRequest.CategoryId
	product.UpdatedAt = time.Now()

	product = service.ProductRepository.Update(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	if product.DeletedAt.Valid {
		panic("product tidak lagi aktif")
	}

	product.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	service.ProductRepository.Delete(ctx, tx, product)
	return helper.ToProductResponse(product)

}
func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)
	if product.DeletedAt.Valid {
		panic("product tidak lagi aktif")
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productsResponse []web.ProductResponse
	for _, product := range products {
		productsResponse = append(productsResponse, helper.ToProductResponse(product))
	}

	return productsResponse
}
