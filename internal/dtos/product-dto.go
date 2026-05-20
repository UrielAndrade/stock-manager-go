package dtos

type CreateProductDTO struct {
	Name             string  `json:"name" binding:"required" example:"Product Name"`
	Price            float64 `json:"price" binding:"required" example:"99.99"`
	Brand            string  `json:"brand" binding:"required" example:"Brand Name"`
	Quantity         int     `json:"quantity" binding:"required" example:"10"`
	IdManufacturerFk int     `json:"manufacturer_id" binding:"required" example:"1"`
}
