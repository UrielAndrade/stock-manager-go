package dtos

type CreateProductDTO struct {
	Name     string  `json:"name" validate:"required" example:"Product Name"`
	Price    float64 `json:"price" validate:"required,gt=0" example:"99.99"`
	BrandID  int     `json:"brand_id" validate:"required" example:"1"`
	Quantity int     `json:"quantity" validate:"required,min=0" example:"10"`
}

type UpdateProductDTO struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	BrandID  int     `json:"brand_id,omitempty"`
	Quantity int     `json:"quantity,omitempty" validate:"omitempty,min=0"`
}

type ProductResponseDTO struct {
	ID       int     `json:"id"`
	BrandID  int     `json:"brand_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
