package dto

type CreateProductRequest struct {
	Name        string  `form:"name" validate:"required"`
	Stock       int     `form:"stock" validate:"required"`
	Price       float64 `form:"price" validate:"required"`
	Description string  `form:"description" validate:"required"`
	Image       string  `form:"image" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string  `form:"name"`
	Stock       int     `form:"stock"`
	Price       float64 `form:"price"`
	Description string  `form:"description"`
	Image       string  `form:"image"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}
