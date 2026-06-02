package dtos

type CreateBrandDTO struct {
	Name           string `json:"name" validate:"required,min=2,max=100" example:"Brand Name"`
	Country        string `json:"country" validate:"required" example:"Brazil"`
	Email          string `json:"email" validate:"required,email" example:"contact@brand.com"`
	FoundationYear int    `json:"foundation_year" validate:"required,gt=1800" example:"1990"`
}

type UpdateBrandDTO struct {
	Name           string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Country        string `json:"country,omitempty"`
	Email          string `json:"email,omitempty" validate:"omitempty,email"`
	FoundationYear int    `json:"foundation_year,omitempty" validate:"omitempty,gt=1800"`
}

type BrandResponseDTO struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	Email          string `json:"email"`
	FoundationYear int    `json:"foundation_year"`
}