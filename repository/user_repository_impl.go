package repository

import (
	"context"
	"database/sql"

	"abdulghofur.me/pshamo-go/model/domain"
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepositoryImpl struct{}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(id, username, phone_number, role, password_hash)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.PhoneNumber, user.Role, user.PasswordHash)
	if err != nil {
		panic(err)
	}

	user = repository.FindById(ctx, tx, user.Id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	user = domain.User{}
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {

}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) domain.User {
	SQL := "select id, username, phone_number, role, password_hash, created_at, updated_at, deleted_at from users where id=?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.PhoneNumber,
			&user.Role,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			panic(err)
		}
	}

	return user
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	var users []domain.User
	return users
}
