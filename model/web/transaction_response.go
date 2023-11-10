package web

import "abdulghofur.me/pshamo-go/model/domain"

type TransactionResponse struct {
	Id            string                   `json:"id"`
	Address       string                   `json:"address"`
	PaymentMethod domain.PaymentMethod     `json:"payment_method"`
	TotalPrice    float32                  `json:"total_price"`
	ShippingPrice float32                  `json:"shipping_price"`
	Status        domain.TransactionStatus `json:"status"`
	CreatedAt     string                   `json:"created_at"`
	UpdatedAt     string                   `json:"updated_at"`
	DeletedAt     string                   `json:"deleted_at"`
}
