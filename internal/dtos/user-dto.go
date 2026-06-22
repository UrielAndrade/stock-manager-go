package dtos

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=6"`
	Birthday string `json:"birthday" validate:"required" example:"2000-01-01"`
	Address  string `json:"address" validate:"required" example:"123 Main St, City, Country"`
	Cpf      string `json:"cpf" validate:"required" example:"123.456.789-00"`
}

type UpdateUserDTO struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
	Birthday string `json:"birthday,omitempty"`
	Address  string `json:"address,omitempty"`
	Cpf      string `json:"cpf,omitempty"`
}

type UserResponseDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Cpf      string `json:"cpf"`
}
