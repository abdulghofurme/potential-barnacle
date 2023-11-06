package domain

import "time"

type User struct {
	Id           string
	Username     string
	PhoneNumber  string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
