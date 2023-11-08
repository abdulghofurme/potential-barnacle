package web

type ProductUpdateRequest struct {
	Id          string  `validate:"required,uuid4"`
	Name        string  `validate:"required"`
	Price       float32 `validate:"required"`
	Description string  `validate:"required"`
	Tags        string  `validate:"required"`
	CategoryId  string  `json:"category_id",validate:"required,uuid4"`
}
