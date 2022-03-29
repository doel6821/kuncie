package service

import (
	"errors"
	"testing"

	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/errormessage"
	repo "github.com/doel6821/kuncie/repository"
	"github.com/stretchr/testify/assert"
)


func TestPromoServiceImpl_GetPromoBySku(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	input := "43N23P"
	res := models.Promo{
		ID:        1,
		Sku:       "43N23P",
		PromoType: "DISCOUNT",
		BonusSku:  "",
		MinQty:    0,
		PayOnly:   0,
		Discount:  10,
	}
	promoRepo.Mock.On("FindPromoBYSku", input).Return(&res, nil)

	result, err := testPromoService.repo.FindPromoBYSku(input)

	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 1)
	assert.Equal(t, result.Sku, "43N23P")
	assert.Equal(t, result.PromoType, "DISCOUNT")
	assert.Equal(t, result.BonusSku, "")
	assert.Equal(t, result.MinQty, 0)
	assert.Equal(t, result.PayOnly, 0)
	assert.Equal(t, result.Discount, 10)
}

func TestPromoServiceImpl_GetPromoBySkuFailed(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	input := "43N2YC"

	promoRepo.Mock.On("FindPromoBYSku", input).Return(nil, errors.New("not found"))

	result, err := testPromoService.repo.FindPromoBYSku(input)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestPromoServiceImpl_CreatePromo(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	input := models.Promo{
		ID:        0,
		Sku:       "43N23P",
		PromoType: "BUY1GET1",
		BonusSku:  "234234",
		MinQty:    0,
		PayOnly:   0,
		Discount:  0,
	}
	res := models.Promo{
		ID:        2,
		Sku:       "234234",
		PromoType: "BUY1GET1",
		BonusSku:  "234234",
		MinQty:    0,
		PayOnly:   0,
		Discount:  0,
	}
	promoRepo.Mock.On("CreatePromo", &input).Return(&res, nil)

	result, err := testPromoService.repo.CreatePromo(&input)

	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 2)
	assert.Equal(t, result.Sku, "234234")
	assert.Equal(t, result.PromoType, "BUY1GET1")
	assert.Equal(t, result.BonusSku, "234234")
	assert.Equal(t, result.MinQty, 0)
	assert.Equal(t, result.PayOnly, 0)
	assert.Equal(t, result.Discount, 0)
}

func TestPromoServiceImpl_CreatePromoFailed(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	input := models.Promo{
		ID:        2,
		Sku:       "",
		PromoType: "BUY1GET1",
		BonusSku:  "234234",
		MinQty:    0,
		PayOnly:   0,
		Discount:  10,
	}

	promoRepo.Mock.On("CreatePromo", &input).Return(nil, errormessage.ErrBadRequest)

	result, err := testPromoService.repo.CreatePromo(&input)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestPromoServiceImpl_UpdatePromo(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	input := models.Promo{
		ID:       2,
		Sku:      "234234",
		PromoType: "BUYXPAYY",
		BonusSku:  "",
		MinQty:    3,
		PayOnly:   2,
		Discount:  0,
	}
	res := models.Promo{
		ID:       2,
		Sku:      "234234",
		PromoType: "BUYXPAYY",
		BonusSku:  "",
		MinQty:    3,
		PayOnly:   2,
		Discount:  0,
	}

	promoRepo.Mock.On("CreatePromo", &input).Return(&res, nil)

	result, err := testPromoService.repo.CreatePromo(&input)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, 2)
	assert.Equal(t, result.Sku, "234234")
	assert.Equal(t, result.PromoType, "BUYXPAYY")
	assert.Equal(t, result.BonusSku, "")
	assert.Equal(t, result.MinQty, 3)
	assert.Equal(t, result.PayOnly, 2)
	assert.Equal(t, result.Discount, 0)
}

func TestPromoServiceImpl_UpdatePromoFailed(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	input := models.Promo{
		ID:       1,
		Sku:      "WQR",
		PromoType: "BUYXPAYY",
		BonusSku:  "",
		MinQty:    3,
		PayOnly:   2,
		Discount:  0,
	}
	

	promoRepo.Mock.On("CreatePromo", &input).Return(nil, errormessage.ErrProductNotExist)

	result, err := testPromoService.repo.CreatePromo(&input)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	
}

func TestPromoServiceImpl_GetAllPromo(t *testing.T) {
	var promoRepo = new(repo.PromoRepositoryMock)
	var testPromoService = NewPromoServiceImpl(promoRepo)
	res := []models.Promo{{
			ID:        1,
			Sku:       "43N23P",
			PromoType: "BUY1GET1",
			BonusSku:  "234234",
			MinQty:    0,
			PayOnly:   0,
			Discount:  0,
		},
		{
			ID:       2,
			Sku:      "234234",
			PromoType: "BUYXPAYY",
			BonusSku:  "",
			MinQty:    3,
			PayOnly:   2,
			Discount:  0,
		},
	}
	promoRepo.Mock.On("FindAllPromo").Return(res, nil)

	result, err := testPromoService.repo.FindAllPromo()

	assert.Equal(t, err, nil)
	assert.Equal(t, result[0].ID, 1)
	assert.Equal(t, result[0].Sku, "43N23P")
	assert.Equal(t, result[0].PromoType, "BUY1GET1")
	assert.Equal(t, result[0].BonusSku, "234234")
	assert.Equal(t, result[0].MinQty, 0)
	assert.Equal(t, result[0].PayOnly, 0)
	assert.Equal(t, result[0].Discount, 0)
	assert.Equal(t, result[1].ID, 2)
	assert.Equal(t, result[1].Sku, "234234")
	assert.Equal(t, result[1].PromoType, "BUYXPAYY")
	assert.Equal(t, result[1].BonusSku, "")
	assert.Equal(t, result[1].MinQty, 3)
	assert.Equal(t, result[1].PayOnly, 2)
	assert.Equal(t, result[1].Discount, 0)
}
