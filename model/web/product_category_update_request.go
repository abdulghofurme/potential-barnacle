package web

type ProductCategoryUpdateRequest struct {
	Id   string `validate:"required,uuid4"`
	Name string `validate:"required"`
}
