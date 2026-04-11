package model

type CreateProductRequest struct {
	Name       string  `json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required,gt=0"`
	Stock      int     `json:"stock" validate:"required,min=0"`
	CategoryID int     `json:"category_id" validate:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryID int     `json:"category_id"`
}

type ProductResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryID   int     `json:"category_id,omitempty"`
	CategoryName string  `json:"category_name,omitempty"`
}
