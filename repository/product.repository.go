package repository

import (
	"github.com/doel6821/kuncie/models"
)

// ProductRepositoryInterface product repository interface
type ProductRepositoryInterface interface {
	CreateProduct(req *models.Product) (*models.Product, error)
	UpdateProduct(req *models.Product) (*models.Product, error)
	FindProductBySku( sku string) (*models.Product, error)
	FindAllProduct() ([]models.Product, error)
}