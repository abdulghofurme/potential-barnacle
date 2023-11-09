package web

import "mime/multipart"

type ProductGalleryUpdateRequest struct {
	Id         string                `validate:"required"`
	File       multipart.File        `validate:"required"`
	FileHeader *multipart.FileHeader `validate:"required"`

	ProductId string `json:"product_id",validate:"required"`
}
