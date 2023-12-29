package dto

type AddProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Price int    `json:"price" validate:"gte=0"`
}
