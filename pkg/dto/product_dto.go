package dto

import "microservice/pkg/models"

type AddProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Price int    `json:"price" validate:"gte=0"`
}

// swagger:response productsResponse
type GetProductsResponse struct {
	// in: body
	products []models.Product
}
