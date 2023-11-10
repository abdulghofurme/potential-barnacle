package controller

import (
	"net/http"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/service"
	"github.com/julienschmidt/httprouter"
)

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func (controller *TransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionRequest := web.TransactionCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionRequest)

	transactionResponse := controller.TransactionService.Create(
		request.Context(),
		transactionRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *TransactionControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionRequest := web.TransactionUpdateRequest{}
	helper.ReadFromRequestBody(request, &transactionRequest)

	transactionRequest.Id = params.ByName("id")
	transactionResponse := controller.TransactionService.Update(
		request.Context(),
		transactionRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *TransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("id")
	transactionResponse := controller.TransactionService.Delete(
		request.Context(),
		transactionId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *TransactionControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("id")
	transactionResponse := controller.TransactionService.FindById(
		request.Context(),
		transactionId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)

}

func (controller *TransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionsResponse := controller.TransactionService.FindAll(
		request.Context(),
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transactionsResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)

}
