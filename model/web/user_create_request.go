package web

import "abdulghofur.me/pshamo-go/model/domain"

type UserCreateRequest struct {
	Username    string
	PhoneNumber string
	Role        domain.UserRoles
	Password    string
}
