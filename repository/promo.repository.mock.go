package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/doel6821/kuncie/models"
)
type PromoRepositoryMock struct {
	Mock mock.Mock
}

func (m *PromoRepositoryMock) CreatePromo(req *models.Promo) (*models.Promo, error) {
	args := m.Mock.Called(req)
	var result *models.Promo
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result = args.Get(0).(*models.Promo) 
		return result, nil
	}
}

func (m *PromoRepositoryMock) FindPromoBYSku(req string) (*models.Promo, error) {
	args := m.Mock.Called(req)
	
	var result *models.Promo
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result = args.Get(0).(*models.Promo) 
		return result, nil
	}
}
func (m *PromoRepositoryMock) FindAllPromo() ([]models.Promo, error) {
	args := m.Mock.Called()
	result := args.Get(0)
	return result.([]models.Promo), args.Error(1)
}