package domain

import (
	"database/sql"
	"fmt"
	"time"
)

// type UserRoles int

// const (
// 	UserRoleUser UserRoles = iota
// 	UserRoleAdmin
// )

// func (roleIndex UserRoles) String() string {
// 	return [...]string{"user", "admin"}[roleIndex]
// }

// func (roleIndex UserRoles) EnumIndex() int {
// 	return int(roleIndex)
// }

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

func (user *User) String() string {
	return fmt.Sprintf(`
	id: %v
	username: %v
	phone_number: %v
	role: %v
	password: %v
	created_at: %v
	updated_at: %v
	deleted_at: %v`,
		user.Id,
		user.Username,
		user.PhoneNumber,
		user.Role,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	)
}
