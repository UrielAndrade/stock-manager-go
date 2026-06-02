package handlers

import (
	"strconv"

	"estoque-go/internal/database"
	"estoque-go/internal/dtos"
	"estoque-go/internal/models"

	"github.com/go-fuego/fuego"
)

func toUserResponse(u models.User) dtos.UserResponseDTO {
	return dtos.UserResponseDTO{
		Id:       u.Id,
		Name:     u.Name,
		Email:    u.Email,
		Birthday: u.Birthday,
		Address:  u.Address,
		Cpf:      u.Cpf,
	}
}

func CreateUser(c fuego.ContextWithBody[dtos.CreateUserDTO]) (dtos.UserResponseDTO, error) {
	input, err := c.Body()
	if err != nil {
		return dtos.UserResponseDTO{}, fuego.BadRequestError{Err: err}
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Birthday: input.Birthday,
		Address:  input.Address,
		Cpf:      input.Cpf,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return dtos.UserResponseDTO{}, err
	}
	return toUserResponse(user), nil
}

func GetUsers(c fuego.ContextNoBody) ([]dtos.UserResponseDTO, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var response []dtos.UserResponseDTO
	for _, u := range users {
		response = append(response, toUserResponse(u))
	}
	return response, nil
}

func UpdateUser(c fuego.ContextWithBody[dtos.UpdateUserDTO]) (dtos.UserResponseDTO, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return dtos.UserResponseDTO{}, fuego.NotFoundError{Err: err}
	}

	input, err := c.Body()
	if err != nil {
		return dtos.UserResponseDTO{}, fuego.BadRequestError{Err: err}
	}

	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		return dtos.UserResponseDTO{}, err
	}

	return toUserResponse(user), nil
}

func DeleteUser(c fuego.ContextNoBody) (any, error) {
	id, _ := strconv.Atoi(c.PathParam("id"))

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return nil, err
	}

	return map[string]string{"message": "Deletado com sucesso"}, nil
}
