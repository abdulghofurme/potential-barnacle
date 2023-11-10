package web

import "abdulghofur.me/pshamo-go/model/domain"

type TransactionCreateRequest struct {
	UserId        string                   `validate:"required"`
	Address       string                   `validate:"required"`
	PaymentMethod domain.PaymentMethod     `validate:"required"`
	TotalPrice    float32                  `validate:"required"`
	ShippingPrice float32                  `validate:"required"`
	Status        domain.TransactionStatus `validate:"required"`
}
