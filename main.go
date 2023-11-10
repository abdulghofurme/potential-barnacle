package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"abdulghofur.me/pshamo-go/app"
	"abdulghofur.me/pshamo-go/config"
	"abdulghofur.me/pshamo-go/controller"
	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/repository"
	"abdulghofur.me/pshamo-go/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	storage := app.NewStorage()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	productCategoryRepository := repository.NewProductCategoryRepository()
	productCategoryService := service.NewProductCategoryService(productCategoryRepository, db, validate)
	productCategoryController := controller.NewProductCategoryController(productCategoryService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	productGalleryRepository := repository.NewProductGalleryRepository()
	productGalleryService := service.NewProductGalleryService(productGalleryRepository, db, validate, storage)
	productGalleryController := controller.NewProductGalleryController(productGalleryService)

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, db, validate)
	transactionController := controller.NewTransactionController(transactionService)

	router := httprouter.New()

	router.POST("/api/users", userController.Create)
	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:id", userController.FindById)
	router.PUT("/api/users/:id", userController.Update)
	router.DELETE("/api/users/:id", userController.Delete)

	router.POST("/api/categories", productCategoryController.Create)
	router.GET("/api/categories", productCategoryController.FindAll)
	router.GET("/api/categories/:id", productCategoryController.FindById)
	router.PUT("/api/categories/:id", productCategoryController.Update)
	router.DELETE("/api/categories/:id", productCategoryController.Delete)

	router.POST("/api/products", productController.Create)
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:id", productController.FindById)
	router.PUT("/api/products/:id", productController.Update)
	router.DELETE("/api/products/:id", productController.Delete)

	router.POST("/api/galleries", productGalleryController.Create)
	router.GET("/api/galleries", productGalleryController.FindAll)
	router.GET("/api/galleries/:id", productGalleryController.FindById)
	router.PUT("/api/galleries/:id", productGalleryController.Update)
	router.DELETE("/api/galleries/:id", productGalleryController.Delete)

	router.POST("/api/transactions", transactionController.Create)
	router.GET("/api/transactions", transactionController.FindAll)
	router.GET("/api/transactions/:id", transactionController.FindById)
	router.PUT("/api/transactions/:id", transactionController.Update)
	router.DELETE("/api/transactions/:id", transactionController.Delete)

	server := http.Server{
		Addr:    config.MyEnv.SERVER_ADDRESS,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
