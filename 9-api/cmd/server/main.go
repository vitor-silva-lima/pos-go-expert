package main

import (
	"first-api/internal/application/service"
	"first-api/internal/infra/database/repository"
	"first-api/internal/infra/webserver/handlers"
	"first-api/internal/root"
	"net/http"
)

func main() {
	conn := root.NewDatabaseConnectionAdapter()
	productRepository, err := repository.NewProductRepository(conn)
	if err != nil {
		panic(err)
	}
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/products/all", productHandler.GetProducts)
	http.HandleFunc("/products/create", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
