package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	json.NewDecoder(r.Body).Decode(&product)

	product.ID = storage.NextID
	storage.NextID++

	storage.Products = append(storage.Products, product)

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
	json.NewEncoder(w).Encode(storage.Products)
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

	for i, p := range storage.Products {
		if p.ID == id {
			json.NewDecoder(r.Body).Decode(&storage.Products[i])
			storage.Products[i].ID = id
			json.NewEncoder(w).Encode(storage.Products[i])
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
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

	for i, p := range storage.Products {
		if p.ID == id {
			storage.Products = append(storage.Products[:i], storage.Products[i+1:]...)
			w.Write([]byte("Deletado com sucesso"))
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}
