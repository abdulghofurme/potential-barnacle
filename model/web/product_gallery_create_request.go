package web

import "mime/multipart"

type ProductGalleryCreateRequest struct {
	File       multipart.File        `validate:"required"`
	FileHeader *multipart.FileHeader `validate:"required"`
	ProductId  string                `json:"product_id",validate:"required"`
}
