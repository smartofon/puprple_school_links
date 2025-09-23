package cart

type ProductCreateRequest struct {
	ProductId   int      `json:"uid" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Price       float64  `json:"price" validate:"required"`
}

type ProductOrder struct {
	ProductId int     `json:"uid" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
}

type OrderCreateResponse struct {
	Products []ProductOrder `json:products`
}

type AnswerTemplate struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}
