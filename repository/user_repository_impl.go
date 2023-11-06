package repository

import (
	"context"
	"database/sql"
	"errors"

	"abdulghofur.me/pshamo-go/model/domain"
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepositoryImpl struct{}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(id, username, phone_number, role, password_hash)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.PhoneNumber, user.Role, user.PasswordHash)
	if err != nil {
		panic(err)
	}

	user, _ = repository.FindById(ctx, tx, user.Id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set username=?, phone_number=?, role=?, password_hash=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.PhoneNumber, user.Role, user.PasswordHash, user.Id)
	if err != nil {
		panic(err)
	}
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "update users set deleted_at=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, user.DeletedAt, user.Id)
	if err != nil {
		panic(err)
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error) {
	SQL := "select id, username, phone_number, role, password_hash, created_at, updated_at, deleted_at from users where id=?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
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

		return user, nil
	} else {
		return user, errors.New("user tidak ditemukan")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, username, phone_number, role, password_hash, created_at, updated_at, deleted_at from users where deleted_at=null"

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	if rows.Next() {
		user := domain.User{}
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
		users = append(users, user)
	}

	return users
}
