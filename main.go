package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"abdulghofur.me/pshamo-go/app"
	"abdulghofur.me/pshamo-go/model/domain"
)

func main() {
	db := app.NewDB()

	rows, err := db.QueryContext(context.Background(), "select id, username, phone_number, created_at, updated_at, deleted_at from users")
	if err != nil {
		panic(err)
	}

	myUUID := uuid.New()

	fmt.Println(myUUID.String())

	var users []domain.User
	if rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		fmt.Println(err)
		users = append(users, user)
	}

	fmt.Println(users)
}
