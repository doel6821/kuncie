package service

import "github.com/doel6821/kuncie/graph/model"

// CheckoutServiceInterface checkout service interface
type CheckoutServiceInterface interface {
	Checkout(req []*model.NewCart) (float64, error)
}