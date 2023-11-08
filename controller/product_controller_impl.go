package controller

import (
	"net/http"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/service"
	"github.com/julienschmidt/httprouter"
)

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := controller.ProductService.Create(
		request.Context(),
		productCreateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productId := params.ByName("id")
	productUpdateRequest.Id = productId

	productResponse := controller.ProductService.Update(
		request.Context(),
		productUpdateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("id")

	productResponse := controller.ProductService.Delete(
		request.Context(),
		productId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("id")

	productResponse := controller.ProductService.FindById(
		request.Context(),
		productId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productsResponse := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productsResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}
