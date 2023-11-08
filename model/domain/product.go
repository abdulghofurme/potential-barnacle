package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          string
	Name        string
	Price       float32
	Description sql.NullString
	Tags        sql.NullString
	CategoryId  string
	Category    ProductCategory
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
