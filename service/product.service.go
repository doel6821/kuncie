package service

import (
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
)

// ProductServiceInterface product service interface
type ProductServiceInterface interface {
	CreateProduct(req *model.NewProduct) (*models.Product, error)
	UpdateProduct(req *model.NewProduct) (*models.Product, error)
	FindProductBySku(sku string) (*models.Product, error)
	FindAllProduct() ([]*models.Product, error)
}