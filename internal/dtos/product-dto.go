package dtos

type CreateProductDTO struct {
	Name             string  `json:"name" validate:"required" example:"Product Name"`
	Price            float64 `json:"price" validate:"required,gt=0" example:"99.99"`
	Brand            string  `json:"brand" validate:"required" example:"Brand Name"`
	Quantity         int     `json:"quantity" validate:"required,min=0" example:"10"`
	IdManufacturerFk int     `json:"manufacturer_id" validate:"required" example:"1"`
}

type UpdateProductDTO struct {
	Name             string  `json:"name,omitempty"`
	Price            float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Brand            string  `json:"brand,omitempty"`
	Quantity         int     `json:"quantity,omitempty" validate:"omitempty,min=0"`
	IdManufacturerFk int     `json:"manufacturer_id,omitempty"`
}

type ProductResponseDTO struct {
	ID               int     `json:"id"`
	IdManufacturerFk int     `json:"manufacturer_id"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	Brand            string  `json:"brand"`
	Quantity         int     `json:"quantity"`
}
