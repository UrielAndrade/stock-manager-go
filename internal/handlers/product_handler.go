package handlers

import (
	"strconv"

	"estoque-go/internal/database"
	"estoque-go/internal/models"

	"github.com/go-fuego/fuego"
)

func CreateProduct(c fuego.ContextWithBody[models.Product]) (models.Product, error) {
	product, err := c.Body()
	if err != nil {
		return models.Product{}, err
	}

	if erro := database.DB.Create(&product).Error; erro != nil {
		return models.Product{}, erro
	}

	return product, nil
}

func GetProducts(c fuego.ContextNoBody) ([]models.Product, error) {
	var products []models.Product
	if erro := database.DB.Find(&products).Error; erro != nil {
		return nil, erro
	}
	return products, nil
}

func UpdateProduct(c fuego.ContextWithBody[models.Product]) (models.Product, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return models.Product{}, fuego.NotFoundError{Err: err}
	}

	input, err := c.Body()
	if err != nil {
		return models.Product{}, fuego.BadRequestError{Err: err}
	}

	if err := database.DB.Model(&product).Updates(input).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func DeleteProduct(c fuego.ContextNoBody) (string, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return "", err
	}

	return "Deletado com sucesso", nil
}
