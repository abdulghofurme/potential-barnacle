package repository

import (
	"context"
	"database/sql"
	"errors"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepositoryImpl struct{}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(id, username, phone_number, role, password_hash) values(?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.PhoneNumber, user.Role, user.PasswordHash)
	helper.PanicIfError(err)

	user, _ = repository.FindById(ctx, tx, user.Id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set username=?, phone_number=?, role=?, password_hash=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.PhoneNumber, user.Role, user.PasswordHash, user.Id)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "update users set deleted_at=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, user.DeletedAt.Time, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error) {
	SQL := "select id, username, phone_number, role, password_hash, created_at, updated_at, deleted_at from users where id=?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
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
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user tidak ditemukan")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, username, phone_number, role, password_hash, created_at, updated_at, deleted_at from users WHERE deleted_at is NULL;"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
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
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindByUsernameAndPhoneNumber(ctx context.Context, tx *sql.Tx, username string, phone_number string) []domain.User {
	SQL := "select id, username, phone_number from users where username=? or phone_number=?"

	rows, err := tx.QueryContext(ctx, SQL, username, phone_number)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	if rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.PhoneNumber,
		)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users

}
