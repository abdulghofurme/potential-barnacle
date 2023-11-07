package domain

import (
	"database/sql"
	"time"
)

type ProductCategory struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
