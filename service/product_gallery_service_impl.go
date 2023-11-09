package service

import (
	"context"
	"database/sql"
	"path/filepath"
	"strings"
	"time"

	"abdulghofur.me/pshamo-go/app"
	"abdulghofur.me/pshamo-go/config"
	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewProductGalleryService(productGalleryRepository repository.ProductGalleryRepository, DB *sql.DB, validate *validator.Validate, storage *app.Storage) ProductGalleryService {
	return &ProductGalleryServiceImpl{
		ProductGalleryRepository: productGalleryRepository,
		DB:                       DB,
		Validate:                 validate,
		Storage:                  storage,
	}
}

type ProductGalleryServiceImpl struct {
	ProductGalleryRepository repository.ProductGalleryRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
	Storage                  *app.Storage
}

func (service *ProductGalleryServiceImpl) Create(ctx context.Context, productGalleryRequest web.ProductGalleryCreateRequest) web.ProductGalleryResponse {
	err := service.Validate.Struct(productGalleryRequest)
	helper.PanicIfErrof(err)

	fileName := productGalleryRequest.FileHeader.Filename
	fileExtension := filepath.Ext(fileName)
	if !strings.Contains(config.ProductGalleryAllowedFormat, fileExtension) {
		panic("file format tidak diizinkan")
	}
	id := uuid.NewString()
	url := service.Storage.Upload(productGalleryRequest.File, fileName, id)

	tx, err := service.DB.Begin()
	helper.PanicIfErrof(err)
	defer helper.CommitOrRollback(tx)

	productGallery := domain.ProductGallery{
		Id:        id,
		Url:       url,
		ProductId: productGalleryRequest.ProductId,
	}

	productGallery = service.ProductGalleryRepository.Create(ctx, tx, productGallery)
	return helper.ToProductGalleryResponse(productGallery)
}

func (service *ProductGalleryServiceImpl) Update(ctx context.Context, productGalleryRequest web.ProductGalleryUpdateRequest) web.ProductGalleryResponse {
	err := service.Validate.Struct(productGalleryRequest)
	helper.PanicIfErrof(err)

	fileName := productGalleryRequest.FileHeader.Filename
	fileExtension := filepath.Ext(fileName)
	if !strings.Contains(config.ProductGalleryAllowedFormat, fileExtension) {
		panic("file format tidak diizinkan")
	}
	url := service.Storage.Upload(productGalleryRequest.File, fileName, productGalleryRequest.Id)

	tx, err := service.DB.Begin()
	helper.PanicIfErrof(err)
	defer helper.CommitOrRollback(tx)

	productGallery, err := service.ProductGalleryRepository.FindById(ctx, tx, productGalleryRequest.Id)
	helper.PanicIfErrof(err)
	if productGallery.DeletedAt.Valid {
		panic("product gallery tidak lagi aktif")
	}

	productGallery.Url = url
	productGallery.ProductId = productGalleryRequest.ProductId
	productGallery.UpdatedAt = time.Now()

	productGallery = service.ProductGalleryRepository.Update(ctx, tx, productGallery)
	return helper.ToProductGalleryResponse(productGallery)
}

func (service *ProductGalleryServiceImpl) Delete(ctx context.Context, productGalleryId string) web.ProductGalleryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErrof(err)
	defer helper.CommitOrRollback(tx)

	productGallery, err := service.ProductGalleryRepository.FindById(ctx, tx, productGalleryId)
	helper.PanicIfErrof(err)
	if productGallery.DeletedAt.Valid {
		panic("product gallery tidak lagi aktif")
	}

	productGallery.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	service.ProductGalleryRepository.Delete(ctx, tx, productGallery)
	return helper.ToProductGalleryResponse(productGallery)
}

func (service *ProductGalleryServiceImpl) FindById(ctx context.Context, productGalleryId string) web.ProductGalleryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErrof(err)
	defer helper.CommitOrRollback(tx)

	productGallery, err := service.ProductGalleryRepository.FindById(ctx, tx, productGalleryId)
	helper.PanicIfErrof(err)
	if productGallery.DeletedAt.Valid {
		panic("product gallery tidak lagi aktif")
	}

	return helper.ToProductGalleryResponse(productGallery)
}

func (service *ProductGalleryServiceImpl) FindAll(ctx context.Context) []web.ProductGalleryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErrof(err)
	defer helper.CommitOrRollback(tx)

	productGalleries := service.ProductGalleryRepository.FindAll(ctx, tx)

	var productGalleriesResponse []web.ProductGalleryResponse
	for _, productGallery := range productGalleries {
		productGalleriesResponse = append(productGalleriesResponse, helper.ToProductGalleryResponse(productGallery))
	}

	return productGalleriesResponse
}
