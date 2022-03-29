package repository

import (
	"github.com/doel6821/kuncie/config"
	"github.com/doel6821/kuncie/models"
)

// CartRepositoryImpl struct implementation
type CartRepositoryImpl struct {

}

// NewCartRepositoryImpl create new cart repository implementation
func NewCartRepositoryImpl() *CartRepositoryImpl {
	return &CartRepositoryImpl{}
}

// CreateCart create new cart
func (c *CartRepositoryImpl) CreateCart(req *models.Cart) (*models.Cart, error) {
	var db = config.SetupDatabaseConnection()
	
	err := db.Save(req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateCartByStatus update cart status
func (c *CartRepositoryImpl) UpdateCartByStatus() (err error) {
	var db = config.SetupDatabaseConnection()
	err = db.Raw("update from carts set status = 1 where status = 0").Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllCart get all cart
func (c *CartRepositoryImpl) GetAllCart() ( []models.Cart,  error) {
	var db = config.SetupDatabaseConnection()
	var carts []models.Cart
	err := db.Where("status = 0").Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}