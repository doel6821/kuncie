package service

import (
	"fmt"

	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/constant"
	"github.com/doel6821/kuncie/package/errormessage"
	repo "github.com/doel6821/kuncie/repository"
)

// CheckoutServiceImpl checkout service implementation
type CheckoutServiceImpl struct {
	productRepo repo.ProductRepositoryInterface
	promoRepo repo.PromoRepositoryInterface
	cartRepo repo.CartRepositoryInterface
}

// NewCheckoutServiceImpl create new checkout service implementation
func NewCheckoutServiceImpl(productRepo repo.ProductRepositoryInterface, promoRepo repo.PromoRepositoryInterface, cartRepo repo.CartRepositoryInterface ) *CheckoutServiceImpl {
	return &CheckoutServiceImpl{
		productRepo,
		promoRepo,
		cartRepo,
	}
}

// Checkout process flow..
func (c *CheckoutServiceImpl) Checkout(req *model.Carts) (float64, error) {
	// change array to map & merge duplicate content
	cartMap := c.ArrToMap(req.Contents)
	
	var pay float64 = 0
	var bonusProductPrice float64 = 0
	

	for key, qtty := range cartMap {
		// check stock avalability & promo
		promo, err := c.promoRepo.FindPromoBYSku(key)
		product, _ := c.productRepo.FindProductBySku(key)
		
		// check available stock
		if product.Quantity < qtty {
			return 0, errormessage.ErrNotEnoughStock
		}
		
		if err == nil && product.Quantity >= qtty {
			
			var productBonus *models.Product
			if promo.BonusSku != "" {
				productBonus, _ = c.productRepo.FindProductBySku(promo.BonusSku)
			}

			// calculate promo depend on promo type
			switch promo.PromoType {
			case constant.BUY1GET1:
				fmt.Println(product.Price, productBonus.Price, qtty)
				bonusProductPrice += (productBonus.Price * float64(qtty))
				pay += (product.Price * float64(qtty))   
				
				// update stock bonus product
				product.Quantity -=  qtty
				c.productRepo.UpdateProduct(product)

			case constant.BUYXPAYY: 
				// Check quantity vs promo minimal quantity
				if qtty >= promo.MinQty {
					// Check multiply
					if qtty / promo.MinQty >= 2 {
						pay += product.Price * float64(promo.PayOnly) *  float64(int(qtty / promo.MinQty)) 
						pay += product.Price * float64(qtty % promo.MinQty)
					} else {
						pay += product.Price * float64(promo.PayOnly)
						pay += product.Price * float64(qtty % promo.MinQty)
					}
				} else {
					pay += float64(qtty) * product.Price
				}

				// update stock bonus product
				product.Quantity -=  qtty
				c.productRepo.UpdateProduct(product)

			case constant.DISCOUNT:
				// Check quantity vs promo minimal quantity
				if qtty >= promo.MinQty {
					pay += float64(qtty) * product.Price * (float64(100-promo.Discount)/100.0)
				} else {
					pay += float64(qtty) * product.Price
				}

				// update stock bonus product
				product.Quantity -=  qtty
				c.productRepo.UpdateProduct(product)
			}
			
		} else {
			// calculate if no Promo
			pay += float64(qtty) * product.Price

			// update stock bonus product
			product.Quantity -=  qtty
			c.productRepo.UpdateProduct(product)
		}
	}
	// update status cart
	c.cartRepo.UpdateCartByStatus()

	fmt.Println(pay - bonusProductPrice)
	return pay - bonusProductPrice, nil

}

// ArrToMap change array to map and merge duplicate
func (c *CheckoutServiceImpl) ArrToMap(req []*model.NewCart) map[string]int {
	res := map[string]int{}

	for _, c:= range req {
		_ , ok := res[c.Sku]; if ok  {
			res[c.Sku] +=  c.Quantity

		} else {
			res[c.Sku] = c.Quantity
		}
		fmt.Println(res, c.Quantity)
	}

	return res
}
