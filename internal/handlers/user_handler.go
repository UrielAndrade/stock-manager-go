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

// CreateProduct godoc
// @Summary Create product
// @Description Create a new product in inventory
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Router /products [post]
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

// GetProducts godoc
// @Summary List products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if erro := database.DB.Find(&users).Error; erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Failure 404 {string} string "not found"
// @Router /products/{id} [put]
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

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {string} string "deleted"
// @Failure 404 {string} string "not found"
// @Router /products/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Deletado com sucesso"))
}
