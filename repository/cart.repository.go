package repository

import "github.com/doel6821/kuncie/models"

// CartRepositoryInterface interface
type CartRepositoryInterface interface {
	CreateCart(req *models.Cart) (*models.Cart, error)
	UpdateCartByStatus() (err error)
	GetAllCart() ([]models.Cart, error)
}