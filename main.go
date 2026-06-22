package main

import (
	"log"

	"estoque-go/internal/database"
	"estoque-go/internal/handlers"
	"estoque-go/internal/models"
	"estoque-go/internal/domain"
	"estoque-go/internal/infra/postgres"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Product{}, &models.Brand{}, &models.User{}, &models.Employee{}, &domain.Order{}, &domain.OrderAudit{})

	s := fuego.NewServer(
		fuego.WithAddr("0.0.0.0:8080"),
	)
	s.OpenAPI.Description().Servers = openapi3.Servers{
		&openapi3.Server{URL: "/", Description: "Local server"},
	}
	s.OpenAPI.Config.DisableDefaultServer = true

	// Product routes
	products := fuego.Group(s, "/products")
	fuego.Get(products, "/", handlers.GetProducts)
	fuego.Post(products, "/", handlers.CreateProduct)
	fuego.Put(products, "/{id}", handlers.UpdateProduct)
	fuego.Delete(products, "/{id}", handlers.DeleteProduct)

	// Order routes
	orders := fuego.Group(s, "/orders")
	orderHandler := handlers.NewOrderHandler(postgres.NewOrderRepository(), postgres.NewAuditRepository())
	fuego.Post(orders, "/", orderHandler.CreateOrder)
	fuego.Get(orders, "/", orderHandler.GetOrders)
	fuego.Get(orders, "/{id}", orderHandler.GetOrder)
	fuego.Put(orders, "/{id}/execute", orderHandler.ExecuteOrder)
	fuego.Put(orders, "/{id}/cancel", orderHandler.CancelOrder)

	// User routes
	users := fuego.Group(s, "/users")
	fuego.Get(users, "/", handlers.GetUsers)
	fuego.Post(users, "/", handlers.CreateUser)
	fuego.Put(users, "/{id}", handlers.UpdateUser)
	fuego.Delete(users, "/{id}", handlers.DeleteUser)

	// Brand routes
	brands := fuego.Group(s, "/brands")
	fuego.Get(brands, "/", handlers.GetBrands)
	fuego.Post(brands, "/", handlers.CreateBrand)
	fuego.Put(brands, "/{id}", handlers.UpdateBrand)
	fuego.Delete(brands, "/{id}", handlers.DeleteBrand)

	log.Println("Server rodando em :8080")
	s.Run()
}
