package handlers

import (
	"encoding/json"
	"net/http"

	"estoque-go/internal/database"
	"estoque-go/internal/models"
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
}
