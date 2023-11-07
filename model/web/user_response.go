package web

import (
	"time"
)

type UserResponse struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   string    `json:"deleted_at"`
}
