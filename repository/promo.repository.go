package repository

import "github.com/doel6821/kuncie/models"

// PromoRepositoryInterface promo repository interface
type PromoRepositoryInterface interface {
	CreatePromo(req *models.Promo) (*models.Promo, error)
	FindPromoBYSku(req string) (*models.Promo, error)
	FindAllPromo() ([]models.Promo, error)
}