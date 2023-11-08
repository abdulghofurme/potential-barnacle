package web

type ProductResponse struct {
	Id          string                  `json:"id"`
	Name        string                  `json:"name"`
	Price       float32                 `json:"price"`
	Description string                  `json:"description"`
	Tags        string                  `json:"tags"`
	Category    ProductCategoryResponse `json:"category"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
	DeletedAt   string                  `json:"deleted_at"`
}
