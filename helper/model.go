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
		CreatedAt: productCategory.CreatedAt,
		UpdatedAt: productCategory.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
