package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"abdulghofur.me/pshamo-go/app"
	"abdulghofur.me/pshamo-go/config"
	"abdulghofur.me/pshamo-go/controller"
	"abdulghofur.me/pshamo-go/repository"
	"abdulghofur.me/pshamo-go/service"
)

func main() {
	db := app.NewDB()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	router.POST("/api/users", userController.Create)
	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:id", userController.FindById)
	router.PUT("/api/users/:id", userController.Update)
	router.DELETE("/api/users/:id", userController.Delete)

	server := http.Server{
		Addr:    config.MyEnv.SERVER_ADDRESS,
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
