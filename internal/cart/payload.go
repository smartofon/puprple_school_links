package cart

type ProductCreateRequest struct {
	ProductId   int      `json:"uid" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Price       float64  `json:"price" validate:"required"`
}
