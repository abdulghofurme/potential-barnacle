package config

import (
	"fmt"
	"os"

	"abdulghofur.me/pshamo-go/helper"
	"github.com/joho/godotenv"
)

type Env struct {
	SERVER_ADDRESS      string
	DB_MYSQL_DSN        string
	STORAGE_BUCKET      string
	STORAGE_FOLDER_NAME string
}

var MyEnv *Env

func init() {
	fmt.Println("config init")
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	MyEnv = &Env{
		SERVER_ADDRESS:      os.Getenv("SERVER_ADDRESS"),
		DB_MYSQL_DSN:        os.Getenv("DB_MYSQL_DSN"),
		STORAGE_BUCKET:      os.Getenv("STORAGE_BUCKET"),
		STORAGE_FOLDER_NAME: os.Getenv("STORAGE_FOLDER_NAME"),
	}
}
