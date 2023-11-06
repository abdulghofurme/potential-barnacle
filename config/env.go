package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_MYSQL_DSN string
}

var MyEnv *Env

func init() {
	fmt.Println("config init")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	MyEnv = &Env{
		DB_MYSQL_DSN: os.Getenv("DB_MYSQL_DSN"),
	}
}
