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
