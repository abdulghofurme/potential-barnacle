package web

import "abdulghofur.me/pshamo-go/model/domain"

type UserUpdateRequest struct {
	Id          string
	Username    string
	PhoneNumber string
	Role        domain.UserRoles
}
