package domain

import (
	"database/sql"
	"time"
)

type UserRoles string

const (
	Undefined     UserRoles = "user"
	UserRoleUser            = "user"
	UserRoleAdmin           = "admin"
)

type User struct {
	Id           string
	Username     string
	PhoneNumber  string
	Role         UserRoles
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}
