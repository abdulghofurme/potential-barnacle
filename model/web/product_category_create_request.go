package web

type ProductCategoryCreateRequest struct {
	Name string `validate:"required"`
}
