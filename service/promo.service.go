package service

import (
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
)

// PromoServiceInterface promo service interface
type PromoServiceInterface interface {
	CreatePromo(req model.NewPromo) (*models.Promo, error)
	UpdatePromo(req model.NewPromo) (*models.Promo, error)
	GetPromoBySku( sku string) (*models.Promo, error)
	GetAllPromo() ([]*models.Promo, error)
}