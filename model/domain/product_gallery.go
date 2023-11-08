package domain

import (
	"database/sql"
	"time"
)

type ProductGallery struct {
	Id        string
	Url       string
	ProductId string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
