package web

import (
	"time"

	"abdulghofur.me/pshamo-go/model/domain"
)

type UserResponse struct {
	Id          string
	Username    string
	PhoneNumber string
	Role        domain.UserRoles
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
