package main

import (
	"log"

	"microservice/internal/api/http/handlers"
	"microservice/internal/api/http/server"
	"microservice/pkg/repositories"
	"microservice/pkg/services"
)

func main() {
	productRepo := repositories.NewProductRepo()
	productService := services.NewProductService(productRepo)
	handler := handlers.NewProductHandler(productService)

	router := server.NewProductRouter(&handler)
	server := server.NewServer(router)

	log.Println(server.InitServer())
}
