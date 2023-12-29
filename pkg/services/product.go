package services

import (
	"microservice/pkg/dto"
	"microservice/pkg/models"
	"microservice/pkg/repositories"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepository repositories.Products
}

func NewProductService(productRepository repositories.Products) ProductService {
	return ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) AddProduct(product dto.AddProductRequest) error {
	newProduct := models.Product{
		Id:    uuid.New().String(),
		Name:  product.Name,
		Type:  product.Type,
		Price: product.Price,
	}
	return p.productRepository.AddProducts(newProduct)
}

func (p *ProductService) GetProducts() []models.Product {
	return p.productRepository.GetProducts()
}
