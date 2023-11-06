package helper

import (
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt.Time,
	}
}
