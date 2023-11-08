package controller

import (
	"encoding/json"
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
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&productCategoryCreateRequest)
	helper.PanicIfErrof(err)

	productCategoryResponse := controller.ProductCategoryService.Create(
		request.Context(),
		productCategoryCreateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   productCategoryResponse,
	}

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfErrof(err)
}

func (controller *ProductCategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoryUpdateRequest := web.ProductCategoryUpdateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&productCategoryUpdateRequest)
	helper.PanicIfErrof(err)

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

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfErrof(err)
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

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfErrof(err)

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

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfErrof(err)
}

func (controller *ProductCategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCategoriesResponse := controller.ProductCategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productCategoriesResponse,
	}

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfErrof(err)
}
