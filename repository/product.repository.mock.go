package repository

import (
	"fmt"

	"github.com/doel6821/kuncie/models"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (m *ProductRepositoryMock) CreateProduct(req *models.Product) (*models.Product, error) {
	args := m.Mock.Called(req)
	fmt.Println(args)
	var result *models.Product
	if args.Get(0) == nil {
		fmt.Println(args.Error(1))
		return nil, args.Error(1)
	} else {
		result = args.Get(0).(*models.Product) 
		return result, nil
	}

}

func (m *ProductRepositoryMock) UpdateProduct(req *models.Product) (*models.Product, error) {
	args := m.Mock.Called(req)
	fmt.Println(args)
	var result *models.Product
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result = args.Get(0).(*models.Product) 
		return result, nil
	}

}

func (m *ProductRepositoryMock)	FindProductBySku( sku string) (*models.Product, error) {
	args := m.Mock.Called(sku)
	fmt.Println(args)
	var result *models.Product
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result = args.Get(0).(*models.Product) 
		return result, nil
	}
	
}

func (m *ProductRepositoryMock)	FindAllProduct() ([]models.Product, error) {
	args := m.Mock.Called()
	result := args.Get(0) 
	return result.([]models.Product), args.Error(1)
}