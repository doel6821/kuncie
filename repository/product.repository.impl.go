package repository

import (
	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/config"
)

// ProductRepositoryImpl struct implementation
type ProductRepositoryImpl struct {

}

// NewProductRepositoryImpl create new product repository implementation
func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

// CreateProduct create new product
func ( repo *ProductRepositoryImpl) CreateProduct(req *models.Product) (*models.Product, error) {
	var db = config.SetupDatabaseConnection()

	err := db.Save(req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateProduct create new product
func ( repo *ProductRepositoryImpl) UpdateProduct(req *models.Product) (*models.Product, error) {
	var db = config.SetupDatabaseConnection()

	err := db.Save(req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FindProductBySku find product by sku
func ( repo *ProductRepositoryImpl) FindProductBySku( sku string) ( *models.Product, error) {
	var db = config.SetupDatabaseConnection()
	var product models.Product
	err := db.Where("sku = ?", sku).Take(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// FindAllProduct find all product
func ( repo *ProductRepositoryImpl) FindAllProduct() ([]models.Product, error) {
	var db = config.SetupDatabaseConnection()
	var products []models.Product
	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}