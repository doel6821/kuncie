package service

import (
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
)

// CartServiceInterface cart service interface
type CartServiceInterface interface {
	CreateCart(req *model.NewCart) (*models.Cart, error) 
	GetAllCart() ([]models.Cart, error)
	
}