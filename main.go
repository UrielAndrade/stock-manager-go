package main

import (
	"log"

	"estoque-go/internal/database"
	"estoque-go/internal/handlers"
	"estoque-go/internal/models"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Product{}, &models.Brand{}, &models.User{}, &models.Employee{})

	s := fuego.NewServer(
		fuego.WithAddr("0.0.0.0:8080"),
	)
	s.OpenAPI.Description().Servers = openapi3.Servers{
		&openapi3.Server{URL: "/", Description: "Local server"},
	}
	s.OpenAPI.Config.DisableDefaultServer = true

	fuego.Get(s, "/products", handlers.GetProducts)
	fuego.Post(s, "/products", handlers.CreateProduct)
	fuego.Put(s, "/products/{id}", handlers.UpdateProduct)
	fuego.Delete(s, "/products/{id}", handlers.DeleteProduct)

	fuego.Get(s, "/users", handlers.GetUsers)
	fuego.Post(s, "/users", handlers.CreateUser)
	fuego.Put(s, "/users/{id}", handlers.UpdateUser)
	fuego.Delete(s, "/users/{id}", handlers.DeleteUser)

	fuego.Get(s, "/brands", handlers.GetBrands)
	fuego.Post(s, "/brands", handlers.CreateBrand)
	fuego.Put(s, "/brands/{id}", handlers.UpdateBrand)
	fuego.Delete(s, "/brands/{id}", handlers.DeleteBrand)

	log.Println("Server rodando em :8080")
	s.Run()
}
