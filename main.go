package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"abdulghofur.me/pshamo-go/config"
)

func main() {
	// db := app.NewDB()

	router := httprouter.New()

	// router.GET("/api.users")

	server := http.Server{
		Addr:    config.MyEnv.SERVER_ADDRESS,
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
