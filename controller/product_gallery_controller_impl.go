package controller

import (
	"net/http"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/service"
	"github.com/julienschmidt/httprouter"
)

func NewProductGalleryController(productGalleryService service.ProductGalleryService) ProductGalleryController {
	return &ProductGalleryControllerImpl{
		ProductGalleryService: productGalleryService,
	}
}

type ProductGalleryControllerImpl struct {
	ProductGalleryService service.ProductGalleryService
}

func (controller *ProductGalleryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productGalleryCreateRequest := web.ProductGalleryCreateRequest{}

	if err := request.ParseMultipartForm(1024); err != nil {
		panic(err)
	}

	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	productGalleryCreateRequest.ProductId = request.FormValue("product_id")
	productGalleryCreateRequest.File = file
	productGalleryCreateRequest.FileHeader = fileHeader

	productGalleryResponse := controller.ProductGalleryService.Create(
		request.Context(),
		productGalleryCreateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   productGalleryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductGalleryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productGalleryUpdateRequest := web.ProductGalleryUpdateRequest{}

	if err := request.ParseMultipartForm(1024); err != nil {
		panic(err)
	}

	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	productGalleryUpdateRequest.ProductId = request.FormValue("product_id")
	productGalleryUpdateRequest.File = file
	productGalleryUpdateRequest.FileHeader = fileHeader

	productGalleryId := params.ByName("id")
	productGalleryUpdateRequest.Id = productGalleryId

	productGalleryResponse := controller.ProductGalleryService.Update(
		request.Context(),
		productGalleryUpdateRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productGalleryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductGalleryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productGalleryId := params.ByName("id")

	productGalleryResponse := controller.ProductGalleryService.Delete(
		request.Context(),
		productGalleryId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productGalleryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductGalleryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productGalleryId := params.ByName("id")

	productGalleryResponse := controller.ProductGalleryService.FindById(
		request.Context(),
		productGalleryId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productGalleryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *ProductGalleryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productGalleriesResponse := controller.ProductGalleryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productGalleriesResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}
