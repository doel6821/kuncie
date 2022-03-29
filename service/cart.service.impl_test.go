package service

import (
	"testing"
	"errors"
	"github.com/doel6821/kuncie/models"
	repo "github.com/doel6821/kuncie/repository"
	"github.com/stretchr/testify/assert"
)


func TestCartServiceImpl_CreateCartSuccess(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var cartRepo = new(repo.CartRepositoryMock)
	var testCartService = NewCartServiceImpl(cartRepo, productRepo)
	input := models.Cart{
		Sku:      "43N23P",
		Quantity: 5,
	}
	res := models.Cart{
		ID:       1,
		Sku:      "43N23P",
		Quantity: 5,
		Status:   0,
	}
	cartRepo.Mock.On("CreateCart", &input).Return(&res, nil)

	result, err := testCartService.repo.CreateCart(&input)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 1)
	assert.Equal(t, result.Sku, "43N23P")
	assert.Equal(t, result.Quantity, 5)
	assert.Equal(t, result.Status, 0)
}

func TestCartServiceImpl_CreateCartFailed(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var cartRepo = new(repo.CartRepositoryMock)
	var testCartService = NewCartServiceImpl(cartRepo, productRepo)

	input := models.Cart{
		Sku:      "",
		Quantity: 5,
	}
	
	cartRepo.Mock.On("CreateCart", &input).Return(nil, errors.New("bad request"))

	result, err := testCartService.repo.CreateCart(&input)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestCartServiceImpl_GetAllCart(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var cartRepo = new(repo.CartRepositoryMock)
	var testCartService = NewCartServiceImpl(cartRepo, productRepo)
	
	res := []models.Cart{{
			ID:        1,
			Sku:       "43N23P",
			Quantity:  1,
		},
		{
			ID:       2,
			Sku:      "234234",
			Quantity:  1,
		},
	}
	cartRepo.Mock.On("GetAllCart").Return(res, nil)

	result, err := testCartService.repo.GetAllCart()

	assert.Equal(t, err, nil)
	assert.Equal(t, result[0].ID, 1)
	assert.Equal(t, result[0].Sku, "43N23P")
	assert.Equal(t, result[0].Quantity, 1)
	assert.Equal(t, result[1].ID, 2)
	assert.Equal(t, result[1].Sku, "234234")
	assert.Equal(t, result[1].Quantity, 1)
}
