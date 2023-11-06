package domain

import (
	"database/sql"
	"time"
)

type User struct {
	Id           string
	Username     string
	PhoneNumber  string
	Role         string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}
