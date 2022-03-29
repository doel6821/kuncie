package repository

import (
	"github.com/doel6821/kuncie/config"
	"github.com/doel6821/kuncie/models"
)

// PromoRepositoryImpl struct implementation
type PromoRepositoryImpl struct {

}

// NewPromoRepositoryImpl create new promo repository implementation
func NewPromoRepositoryImpl() *PromoRepositoryImpl {
	return &PromoRepositoryImpl{}
}

// CreatePromo create new promo
func(p *PromoRepositoryImpl) CreatePromo(req *models.Promo) (*models.Promo, error) {
	var db = config.SetupDatabaseConnection()
	
	err := db.Save(req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FindPromoBYSku find promo bu sku
func(p *PromoRepositoryImpl) FindPromoBYSku(sku string) ( *models.Promo,  error) {
	var db = config.SetupDatabaseConnection()
	var promo models.Promo
	err := db.Where("sku = ?", sku).Take(&promo).Error
	if err != nil {
		return nil, err
	}
	return &promo, nil
}

// FindAllPromo Find all promo
func(p *PromoRepositoryImpl) FindAllPromo() ([]models.Promo, error) {
	var db = config.SetupDatabaseConnection()
	var listpromo []models.Promo
	err := db.Find(&listpromo).Error
	if err != nil {
		return nil, err
	}
	
	return listpromo, nil
}