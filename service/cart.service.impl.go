package service

import (
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/errormessage"
	repo "github.com/doel6821/kuncie/repository"
)

// CartServiceImpl cart service implementation
type CartServiceImpl struct {
	repo repo.CartRepositoryInterface
	productRepo repo.ProductRepositoryInterface
}

// NewCartServiceImpl create new cart service implementation
func NewCartServiceImpl( repo repo.CartRepositoryInterface, productRepo repo.ProductRepositoryInterface) *CartServiceImpl {
	return &CartServiceImpl{
		repo,
		productRepo,
	}
}

// CreateCart service create cart
func (c *CartServiceImpl) CreateCart( newCart *model.NewCart) (*models.Cart, error) {
	if newCart == nil {
		return nil, errormessage.ErrBadRequest
	}
	cart, err := models.NewCart(newCart.Quantity, newCart.Sku) 
	if err != nil {
		return nil, errormessage.ErrBadRequest
	}
	_ , err = c.productRepo.FindProductBySku(newCart.Sku)
	if err != nil {
		return nil, errormessage.ErrProductNotExist
	}

	cart, err = c.repo.CreateCart(cart)
	if err != nil {
		return nil, errormessage.ErrSaveDb
	}

	return cart, nil
}

// GetAllCart service get all cart
func (c *CartServiceImpl) GetAllCart() ([]models.Cart, error) {
	carts, err := c.repo.GetAllCart()
	if err != nil {
		return nil, errormessage.ErrNotFound
	}

	return carts, nil
}



