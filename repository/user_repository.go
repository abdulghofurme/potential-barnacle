package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
