package dtos

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=6"`
	Birthday string `json:"birthday" binding:"required" example:"2000-01-01"`
	Address  string `json:"address" binding:"required" example:"123 Main St, City, Country"`
	Cpf      string `json:"cpf" binding:"required" example:"123.456.789-00"`
}
