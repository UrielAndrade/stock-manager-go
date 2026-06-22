package handlers

import (
	"strconv"

	"estoque-go/internal/database"
	"estoque-go/internal/dtos"
	"estoque-go/internal/models"

	"github.com/go-fuego/fuego"
)

func toBrandResponse(b models.Brand) dtos.BrandResponseDTO {
	return dtos.BrandResponseDTO{
		ID:             b.ID,
		Name:           b.Name,
		Country:        b.Country,
		Email:          b.Email,
		FoundationYear: b.FoundationYear,
	}
}

func CreateBrand(c fuego.ContextWithBody[dtos.CreateBrandDTO]) (dtos.BrandResponseDTO, error) {
	input, err := c.Body()
	if err != nil {
		return dtos.BrandResponseDTO{}, fuego.BadRequestError{Err: err}
	}

	// Verificar se o e-mail já existe para evitar erro de constraint no banco e pulo de ID
	var count int64
	database.DB.Model(&models.Brand{}).Where("email = ?", input.Email).Count(&count)
	if count > 0 {
		return dtos.BrandResponseDTO{}, fuego.ConflictError{
			Title:  "Conflito de cadastro",
			Detail: "Uma marca com este e-mail já está cadastrada",
		}
	}

	brand := models.Brand{
		Name:           input.Name,
		Country:        input.Country,
		Email:          input.Email,
		FoundationYear: input.FoundationYear,
	}

	if err := database.DB.Create(&brand).Error; err != nil {
		return dtos.BrandResponseDTO{}, err
	}

	return toBrandResponse(brand), nil
}

func GetBrands(c fuego.ContextNoBody) ([]dtos.BrandResponseDTO, error) {
	var brands []models.Brand
	if err := database.DB.Find(&brands).Error; err != nil {
		return nil, err
	}

	var response []dtos.BrandResponseDTO
	for _, b := range brands {
		response = append(response, toBrandResponse(b))
	}
	return response, nil
}

func UpdateBrand(c fuego.ContextWithBody[dtos.UpdateBrandDTO]) (dtos.BrandResponseDTO, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	var brand models.Brand
	if err := database.DB.First(&brand, id).Error; err != nil {
		return dtos.BrandResponseDTO{}, fuego.NotFoundError{Err: err}
	}

	input, err := c.Body()
	if err != nil {
		return dtos.BrandResponseDTO{}, fuego.BadRequestError{Err: err}
	}

	if err := database.DB.Model(&brand).Updates(input).Error; err != nil {
		return dtos.BrandResponseDTO{}, err
	}

	return toBrandResponse(brand), nil
}

func DeleteBrand(c fuego.ContextNoBody) (any, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	if err := database.DB.Delete(&models.Brand{}, id).Error; err != nil {
		return nil, err
	}

	return map[string]string{"message": "Deletado com sucesso"}, nil
}
