package repositories

import (
	"microservice/pkg/models"

	"github.com/google/uuid"
)

type Products struct {
	products []models.Product
}

func NewProductRepo() Products {
	return Products{
		products: []models.Product{
			{
				Id:    uuid.New().String(),
				Name:  "TV",
				Type:  "Tech",
				Price: 500,
			},
			{
				Id:    uuid.New().String(),
				Name:  "Laptop",
				Type:  "Tech",
				Price: 1000,
			},
			{
				Id:    uuid.New().String(),
				Name:  "Coffee Maker",
				Type:  "Appliance",
				Price: 150,
			},
		},
	}
}

func (p *Products) GetProducts() []models.Product {
	return p.products
}

func (p *Products) AddProduct(product models.Product) error {
	p.products = append(p.products, product)
	return nil
}
