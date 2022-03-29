package service

import (
	"errors"
	"testing"

	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/errormessage"
	repo "github.com/doel6821/kuncie/repository"
	"github.com/stretchr/testify/assert"
)


func TestProductServiceImpl_FindProductBySku(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	input := "43N23P"
	res := models.Product{
		ID:       1,
		Sku:      "43N23P",
		Name:     "Macbook Pro",
		Price:    5399.99,
		Quantity: 5,
	}
	productRepo.Mock.On("FindProductBySku", input).Return(&res)

	result, err := testProductService.repo.FindProductBySku(input)

	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 1)
	assert.Equal(t, result.Sku, "43N23P")
	assert.Equal(t, result.Name, "Macbook Pro")
	assert.Equal(t, result.Quantity, 5)
	assert.Equal(t, result.Price, 5399.99)
}

func TestProductServiceImpl_FindProductBySkuFailed(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	input := "43N23P"

	productRepo.Mock.On("FindProductBySku", input).Return(nil, errors.New("not found"))

	result, err := testProductService.repo.FindProductBySku(input)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	input := models.Product{
		ID:       0,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.99,
		Quantity: 10,
	}
	res := models.Product{
		ID:       2,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.99,
		Quantity: 10,
	}
	productRepo.Mock.On("CreateProduct", &input).Return(&res, nil)

	result, err := testProductService.repo.CreateProduct(&input)

	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 2)
	assert.Equal(t, result.Sku, "234234")
	assert.Equal(t, result.Name, "Raspberry Pi B")
	assert.Equal(t, result.Quantity, 10)
	assert.Equal(t, result.Price, 30.99)
}

func TestProductServiceImpl_CreateProductFailed(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	input := models.Product{
		ID:       0,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    0,
		Quantity: 10,
	}

	productRepo.Mock.On("CreateProduct", &input).Return(nil, errormessage.ErrBadRequest)

	result, err := testProductService.repo.CreateProduct(&input)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestProductServiceImpl_UpdateProduct(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	input := models.Product{
		ID:       2,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.99,
		Quantity: 5,
	}
	res := models.Product{
		ID:       2,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.99,
		Quantity: 5,
	}

	productRepo.Mock.On("CreateProduct", &input).Return(&res, nil)

	result, err := testProductService.repo.CreateProduct(&input)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 2)
	assert.Equal(t, result.Sku, "234234")
	assert.Equal(t, result.Name, "Raspberry Pi B")
	assert.Equal(t, result.Quantity, 5)
	assert.Equal(t, result.Price, 30.99)
}

func TestProductServiceImpl_UpdateProductFailed(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	input := models.Product{
		ID:       1,
		Sku:      "WQR",
		Name:     "Raspberry Pi B",
		Price:    30.99,
		Quantity: 5,
	}
	

	productRepo.Mock.On("CreateProduct", &input).Return(nil, errormessage.ErrProductNotExist)

	result, err := testProductService.repo.CreateProduct(&input)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	
}

func TestProductServiceImpl_FindAllProduct(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var testProductService = NewProductServiceImpl(productRepo)	
	res := []models.Product{{
			ID:       1,
			Sku:      "43N23P",
			Name:     "Macbook Pro",
			Price:    5399.99,
			Quantity: 5,
		},
		{
			ID:       2,
			Sku:      "234234",
			Name:     "Raspberry Pi B",
			Price:    30.99,
			Quantity: 10,
		},
	}
	productRepo.Mock.On("FindAllProduct").Return(res, nil)

	result, err := testProductService.repo.FindAllProduct()

	assert.Equal(t, err, nil)
	assert.Equal(t, result[0].ID, 1)
	assert.Equal(t, result[0].Sku, "43N23P")
	assert.Equal(t, result[0].Name, "Macbook Pro")
	assert.Equal(t, result[0].Quantity, 5)
	assert.Equal(t, result[0].Price, 5399.99)
	assert.Equal(t, result[1].ID, 2)
	assert.Equal(t, result[1].Sku, "234234")
	assert.Equal(t, result[1].Name, "Raspberry Pi B")
	assert.Equal(t, result[1].Quantity, 10)
	assert.Equal(t, result[1].Price, 30.99)
}
