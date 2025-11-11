package product

type ProductCreateRequest struct {
	Name        string  `json:"name" validate:"required,min=10,max=100"`
	Description string  `json:"description" validate:"omitempty,max=1000"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Quantity    int     `json:"quantity" validate:"required,min=0"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=10,max=100"`
	Description string  `json:"description" validate:"omitempty,max=1000"`
	Price       float64 `json:"price" validate:"omitempty,min=1"`
	Quantity    int     `json:"quantity" validate:"omitempty,min=0"`
}
