package service

import (
	"context"

	"abdulghofur.me/pshamo-go/model/web"
)

type UserService interface {
	Create(ctx context.Context, userRequest web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, userRequest web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId string) web.UserResponse
	FindById(ctx context.Context, userId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
