package handlers

import (
	"encoding/json"
	"first-api/internal/application/service"
	"first-api/internal/dto"
	"net/http"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productDtoInput := &dto.CreateProductDtoInput{}
	err := json.NewDecoder(r.Body).Decode(productDtoInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ph.ProductService.CreateProduct(productDtoInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
