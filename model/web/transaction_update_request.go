package web

import "abdulghofur.me/pshamo-go/model/domain"

type TransactionUpdateRequest struct {
	Id            string                   `validate:"required"`
	Address       string                   `validate:"required"`
	PaymentMethod domain.PaymentMethod     `validate:"required"`
	Status        domain.TransactionStatus `validate:"required"`
}
