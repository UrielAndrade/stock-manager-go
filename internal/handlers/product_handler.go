package handlers

import (
	"strconv"

	"estoque-go/internal/database"
	"estoque-go/internal/dtos"
	"estoque-go/internal/models"

	"github.com/go-fuego/fuego"
)

func toProductResponse(p models.Product) dtos.ProductResponseDTO {
	return dtos.ProductResponseDTO{
		ID:               p.ID,
		IdManufacturerFk: p.IdManufacturerFk,
		Name:             p.Name,
		Price:            p.Price,
		Brand:            p.Brand,
		Quantity:         p.Quantity,
	}
}

func CreateProduct(c fuego.ContextWithBody[dtos.CreateProductDTO]) (dtos.ProductResponseDTO, error) {
	input, err := c.Body()
	if err != nil {
		return dtos.ProductResponseDTO{}, fuego.BadRequestError{Err: err}
	}

	product := models.Product{
		IdManufacturerFk: input.IdManufacturerFk,
		Name:             input.Name,
		Price:            input.Price,
		Brand:            input.Brand,
		Quantity:         input.Quantity,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return dtos.ProductResponseDTO{}, err
	}

	return toProductResponse(product), nil
}

func GetProducts(c fuego.ContextNoBody) ([]dtos.ProductResponseDTO, error) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	var response []dtos.ProductResponseDTO
	for _, p := range products {
		response = append(response, toProductResponse(p))
	}
	return response, nil
}

func UpdateProduct(c fuego.ContextWithBody[dtos.UpdateProductDTO]) (dtos.ProductResponseDTO, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return dtos.ProductResponseDTO{}, fuego.NotFoundError{Err: err}
	}

	input, err := c.Body()
	if err != nil {
		return dtos.ProductResponseDTO{}, fuego.BadRequestError{Err: err}
	}

	if err := database.DB.Model(&product).Updates(input).Error; err != nil {
		return dtos.ProductResponseDTO{}, err
	}

	return toProductResponse(product), nil
}

func DeleteProduct(c fuego.ContextNoBody) (any, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return nil, err
	}

	return map[string]string{"message": "Deletado com sucesso"}, nil
}
