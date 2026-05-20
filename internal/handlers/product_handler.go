package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"estoque-go/internal/database"
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
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	if erro := database.DB.Create(&product).Error; erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	if erro := database.DB.Create(&product).Error; erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// GetProducts godoc
// @Summary List products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if erro := database.DB.Find(&products).Error; erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
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
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	var input models.Product
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Model(&product).Updates(input).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
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
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Deletado com sucesso"))
}
