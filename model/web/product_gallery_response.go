package web

type ProductGalleryResponse struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	ProductId string `json:"product_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
