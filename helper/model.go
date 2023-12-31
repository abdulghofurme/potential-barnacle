package helper

import (
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	deletedAt := user.DeletedAt.Time.String()
	if !user.DeletedAt.Valid {
		deletedAt = ""
	}

	return web.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func ToProductCategoryResponse(productCategory domain.ProductCategory) web.ProductCategoryResponse {
	deletedAt := productCategory.DeletedAt.Time.String()
	if !productCategory.DeletedAt.Valid {
		deletedAt = ""
	}

	return web.ProductCategoryResponse{
		Id:        productCategory.Id,
		Name:      productCategory.Name,
		CreatedAt: productCategory.CreatedAt.String(),
		UpdatedAt: productCategory.UpdatedAt.String(),
		DeletedAt: deletedAt,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	deletedAt := product.DeletedAt.Time.String()
	if !product.DeletedAt.Valid {
		deletedAt = ""
	}

	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description.String,
		Tags:        product.Tags.String,
		Category: web.ProductCategoryResponse{
			Id:   product.Category.Id,
			Name: product.Category.Name,
		},
		CreatedAt: product.CreatedAt.String(),
		UpdatedAt: product.UpdatedAt.String(),
		DeletedAt: deletedAt,
	}
}

func ToProductGalleryResponse(productGallery domain.ProductGallery) web.ProductGalleryResponse {
	deletedAt := productGallery.DeletedAt.Time.String()
	if !productGallery.DeletedAt.Valid {
		deletedAt = ""
	}

	return web.ProductGalleryResponse{
		Id:        productGallery.Id,
		Url:       productGallery.Url,
		ProductId: productGallery.ProductId,
		CreatedAt: productGallery.CreatedAt.String(),
		UpdatedAt: productGallery.UpdatedAt.String(),
		DeletedAt: deletedAt,
	}
}

func ToTransactionResponse(transaction domain.Transaction) web.TransactionResponse {
	deletedAt := transaction.DeletedAt.Time.String()
	if !transaction.DeletedAt.Valid {
		deletedAt = ""
	}

	return web.TransactionResponse{
		Id:            transaction.Id,
		Address:       transaction.Address,
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt.String(),
		UpdatedAt:     transaction.CreatedAt.String(),
		DeletedAt:     deletedAt,
	}
}
