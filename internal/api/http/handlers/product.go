// Package classification Product API.
//
//	Schemes: http
//	Host: localhost
//	BasePath: /
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
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

// GetProduct lists products to the user
func (h *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET /list products listProducts
	//
	// Lists all products.
	// This will show all available products by default.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       200: productsResponse

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
