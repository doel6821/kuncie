package repository

import (

	"github.com/doel6821/kuncie/models"
	"github.com/stretchr/testify/mock"
)
type CartRepositoryMock struct {
	Mock mock.Mock
}


func (m *CartRepositoryMock) CreateCart(req *models.Cart) (*models.Cart, error) {
	args := m.Mock.Called(req)
	
	var result *models.Cart
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result = args.Get(0).(*models.Cart) 
		return result, nil
	}
	
}

func (m *CartRepositoryMock) UpdateCartByStatus() (err error) {
	args := m.Mock.Called()
	return args.Error(0)
}

func (m *CartRepositoryMock) GetAllCart() ([]models.Cart, error) {
	args := m.Mock.Called()
	var result []models.Cart
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result = args.Get(0).([]models.Cart) 
		return result, nil
	}
}