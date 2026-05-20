package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"estoque-go/internal/database"
	"estoque-go/internal/dtos"
	"estoque-go/internal/models"

	"github.com/gorilla/mux"
)

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserDTO true "User data"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var input dtos.CreateUserDTO
	if erro := json.NewDecoder(r.Body).Decode(&input); erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
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
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// GetUsers godoc
// @Summary List users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {string} string "Internal server error"
// @Router /users [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if erro := database.DB.Find(&users).Error; erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 404 {string} string "User not found"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "deleted"
// @Failure 500 {string} string "Internal server error"
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Deletado com sucesso"))
}
