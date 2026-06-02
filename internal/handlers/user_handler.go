package handlers

import (
	"strconv"

	"estoque-go/internal/database"
	"estoque-go/internal/dtos"
	"estoque-go/internal/models"

	"github.com/go-fuego/fuego"
)

func CreateUser(c fuego.ContextWithBody[dtos.CreateUserDTO]) (models.User, error) {
	input, err := c.Body()
	if err != nil {
		return models.User{}, fuego.BadRequestError{Err: err}
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Birthday: input.Birthday,
		Address:  input.Address,
		Cpf:      input.Cpf,
	}

	if erro := database.DB.Create(&user).Error; erro != nil {
		return models.User{}, erro
	}
	return user, nil
}

func GetUser(c fuego.ContextNoBody) ([]models.User, error) {
	var users []models.User
	if erro := database.DB.Find(&users).Error; erro != nil {
		return nil, erro
	}
	return users, nil
}

func UpdateUser(c fuego.ContextWithBody[models.User]) (models.User, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, fuego.NotFoundError{Err: err}
	}

	input, err := c.Body()
	if err != nil {
		return models.User{}, fuego.BadRequestError{Err: err}
	}

	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(c fuego.ContextNoBody) (string, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return "", err
	}

	return "Deletado com sucesso", nil
}
