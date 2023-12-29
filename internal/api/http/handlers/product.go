package handlers

import (
	"encoding/json"
	"microservice/pkg/dto"
	"microservice/pkg/services"

	"net/http"
)

type ProductHandler struct {
	productSvc services.ProductService
}

func NewProductHandler(productSvc services.ProductService) ProductHandler {
	return ProductHandler{
		productSvc: productSvc,
	}
}

func (h *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products, err := json.Marshal(h.productSvc.GetProducts())
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(products))
}

func (h *ProductHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(dto.AddProductRequest{}).(dto.AddProductRequest)

	h.productSvc.AddProduct(product)
	rw.Write([]byte("Add Successfull"))
}
