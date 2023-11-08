package app

import (
	"database/sql"
	"time"

	"abdulghofur.me/pshamo-go/config"
	"abdulghofur.me/pshamo-go/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", config.MyEnv.DB_MYSQL_DSN)
	helper.PanicIfErrof(err)

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
