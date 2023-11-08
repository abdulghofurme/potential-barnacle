package web

type ProductCreateRequest struct {
	Name        string  `validate:"required"`
	Price       float32 `validate:"required"`
	Description string  `validate:"required"`
	Tags        string  `validate:"required"`
	CategoryId  string  `json:"category_id",validate:"required,uuid4"`
}
