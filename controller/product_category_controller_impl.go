package controller

import (
	"net/http"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/service"
	"github.com/julienschmidt/httprouter"
)

func NewProductCategoryController(productCategoryService service.ProductCategoryService) ProductCategoryController {
	return &ProductCategoryControllerImpl{
		ProductCategoryService: productCategoryService,
	}
}

type ProductCategoryControllerImpl struct {
	ProductCategoryService service.ProductCategoryService
}

func (controller *ProductCategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoryCreateRequest := web.ProductCategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &productCategoryCreateRequest)

	productCategoryResponse := controller.ProductCategoryService.Create(
		request.Context(),
		productCategoryCreateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   productCategoryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductCategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoryUpdateRequest := web.ProductCategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &productCategoryUpdateRequest)

	productCategoryId := params.ByName("id")
	productCategoryUpdateRequest.Id = productCategoryId

	productCategoryResponse := controller.ProductCategoryService.Update(
		request.Context(),
		productCategoryUpdateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productCategoryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductCategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoryId := params.ByName("id")

	productCategoryResponse := controller.ProductCategoryService.Delete(
		request.Context(),
		productCategoryId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productCategoryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductCategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoryId := params.ByName("id")

	productCategoryResponse := controller.ProductCategoryService.FindById(
		request.Context(),
		productCategoryId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productCategoryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductCategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoriesResponse := controller.ProductCategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productCategoriesResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}
