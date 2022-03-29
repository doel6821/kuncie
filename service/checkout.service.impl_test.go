package service

import (
	"testing"
	repo "github.com/doel6821/kuncie/repository"
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/errormessage"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutServiceImpl_Checkout1(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var promoRepo = new(repo.PromoRepositoryMock)
	var cartRepo = new(repo.CartRepositoryMock)
	var testCheckoutService = NewCheckoutServiceImpl(productRepo,promoRepo,cartRepo)

	input := model.Carts{
		Contents: []*model.NewCart{
			{
				Sku:      "43N23P",
				Quantity: 1,
			},
			{
				Sku:      "234234",
				Quantity: 1,
			},
		},
	}

	pr1 := &models.Promo{
		ID:       2,
		Sku:      "43N23P",
		PromoType: "BUY1GET1",
		BonusSku:  "234234",
		MinQty:    0,
		PayOnly:   0,
		Discount:  0,
	}
	product1 := &models.Product{
		ID:       1,
		Sku:      "43N23P",
		Name:     "Macbook Pro",
		Price:    5399.99,
		Quantity: 5,
	}

	product2 := &models.Product{
		ID:       2,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.00,
		Quantity: 10,
	}

	product1a := &models.Product{
		ID:       1,
		Sku:      "43N23P",
		Name:     "Macbook Pro",
		Price:    5399.99,
		Quantity: 4,
	}

	product2a := &models.Product{
		ID:       2,
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.00,
		Quantity: 9,
	}

	promoRepo.Mock.On("FindPromoBYSku", "43N23P").Return(pr1, nil)
	promoRepo.Mock.On("FindPromoBYSku", "234234").Return(nil, errormessage.ErrNotFound)
	productRepo.Mock.On("FindProductBySku","43N23P").Return(product1, nil)
	productRepo.Mock.On("FindProductBySku","234234").Return(product2, nil)
	productRepo.Mock.On("UpdateProduct",product1a).Return(product1a, nil)
	productRepo.Mock.On("UpdateProduct",product2a).Return(product2a, nil)
	cartRepo.Mock.On("UpdateCartByStatus").Return(nil)
	
	res, err := testCheckoutService.Checkout(&input)
	assert.Equal(t, 5399.99, res )
	assert.Nil(t, err)
}

func TestCheckoutServiceImpl_Checkout2(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var promoRepo = new(repo.PromoRepositoryMock)
	var cartRepo = new(repo.CartRepositoryMock)
	var testCheckoutService = NewCheckoutServiceImpl(productRepo,promoRepo,cartRepo)

	input := model.Carts{
		Contents: []*model.NewCart{
			{
				Sku:      "120P90",
				Quantity: 3,
			},
		},
	}

	pr1 := &models.Promo{
		ID:       2,
		Sku:      "120P90",
		PromoType: "BUYXPAYY",
		BonusSku:  "",
		MinQty:    3,
		PayOnly:   2,
		Discount:  0,
	}

	product := &models.Product{
		ID:       2,
		Sku:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 10,
	}

	product1a := &models.Product{
		ID:       2,
		Sku:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 7,
	}


	promoRepo.Mock.On("FindPromoBYSku", "120P90").Return(pr1, nil)
	productRepo.Mock.On("FindProductBySku","120P90").Return(product, nil)
	productRepo.Mock.On("UpdateProduct",product1a).Return(product1a, nil)
	cartRepo.Mock.On("UpdateCartByStatus").Return(nil)
	
	res, err := testCheckoutService.Checkout(&input)
	assert.Equal(t, 99.98, res )
	assert.Nil(t, err)
}

func TestCheckoutServiceImpl_Checkout3(t *testing.T) {
	var productRepo = new(repo.ProductRepositoryMock)
	var promoRepo = new(repo.PromoRepositoryMock)
	var cartRepo = new(repo.CartRepositoryMock)
	var testCheckoutService = NewCheckoutServiceImpl(productRepo,promoRepo,cartRepo)

	input := model.Carts{
		Contents: []*model.NewCart{
			{
				Sku:      "A304SD",
				Quantity: 1,
			},
			{
				Sku:      "A304SD",
				Quantity: 1,
			},
			{
				Sku:      "A304SD",
				Quantity: 1,
			},
		},
	}

	pr1 := &models.Promo{
		ID:       2,
		Sku:      "A304SD",
		PromoType: "DISCOUNT",
		BonusSku:  "",
		MinQty:    0,
		PayOnly:   0,
		Discount:  10,
	}
	product1 := &models.Product{
		ID:       3,
		Sku:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.50,
		Quantity: 10,
	}

	product1a := &models.Product{
		ID:       3,
		Sku:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.50,
		Quantity: 7,
	}

	promoRepo.Mock.On("FindPromoBYSku", "A304SD").Return(pr1, nil)
	productRepo.Mock.On("FindProductBySku","A304SD").Return(product1, nil)
	productRepo.Mock.On("UpdateProduct",product1a).Return(product1a, nil)
	cartRepo.Mock.On("UpdateCartByStatus").Return(nil)
	
	res, err := testCheckoutService.Checkout(&input)
	assert.Equal(t , 295.65000000000003, res )
	assert.Nil(t, err)
}
