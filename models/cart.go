package models

import (
	"github.com/doel6821/kuncie/package/errormessage"
)

// Cart is represent table carts
type Cart struct {
	ID       int     `json:"id"`
	Sku      string  `json:"sku"`
	Quantity int     `json:"quantity"`
	Status   int     `json:"status"`
}

// TableName cart table name 
func (t *Cart) TableName() string {
	return "carts"
}

// NewCart generate new cart
func NewCart(quantity int, sku string) (*Cart, error) {
	cart := &Cart{
		ID:       0,
		Sku:      sku,
		Quantity: quantity,
		Status:   0,
	}
	err := cart.validate()
	if err != nil {
		return nil, err
	}

	return cart, nil
}

// validate create cart validation
func (c *Cart) validate() error {
	if c.Sku == "" || c.Quantity == 0 {
		return errormessage.ErrBadRequest
	}
	return nil
}
