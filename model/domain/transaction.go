package domain

import (
	"database/sql"
	"time"
)

type PaymentMethod string

const (
	PaymentManual PaymentMethod = "manual"
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "pending"
)

type Transaction struct {
	Id            string
	UserId        string
	Address       string
	PaymentMethod PaymentMethod
	TotalPrice    float32
	ShippingPrice float32
	Status        TransactionStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}
