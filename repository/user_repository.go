package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, id string)
	FindById(ctx context.Context, tx *sql.Tx, id string) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
