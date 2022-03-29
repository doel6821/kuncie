package models

import (
	"github.com/doel6821/kuncie/package/constant"

	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/package/errormessage"
)

// Promo is represent table promo
type Promo struct {
	ID           int     `json:"id"`
	Sku          string  `json:"sku"`
	PromoType    string  `json:"promoType"`
	BonusSku     string  `json:"bonusSku"`
	MinQty       int     `json:"minQty"`
	PayOnly      int     `json:"payOnly"`
	Discount     int     `json:"discount"`
}

// TableName promo table name 
func (t *Promo) TableName() string {
	return "promo"
}

// PromoDetail detail product promotion
type PromoDetail struct {
	BonusSku string `json:"bonusSku"`
	MinQty   int    `json:"minQty"`
	PayOnly  int    `json:"payOnly"`
	Discount int    `json:"discount"`
}

// NewPromo generate new promo
func NewPromo(sku, code string, detail *model.PromoDetail) (*Promo, error) {
	promo := &Promo{
		ID:           0,
		Sku:          sku,
		PromoType:    code,
		BonusSku:     *detail.Bonussku,
		MinQty:       *detail.Minqty,
		PayOnly:      *detail.Payonly,
		Discount:     *detail.Discount,
	}

	err := promo.validate()
	if err != nil {
		return nil , err
	}

	return promo, nil
}

// validate create promo validation
func (p *Promo) validate() error {
	if p.Sku == "" && p.PromoType == "" {
		return errormessage.ErrBadRequest
	}

	if !(p.PromoType == constant.BUY1GET1 || p.PromoType == constant.BUYXPAYY || p.PromoType == constant.DISCOUNT) {
		return errormessage.ErrBadRequestPromoType
	} 

	switch p.PromoType {
	case constant.BUY1GET1:
		if p.BonusSku == "" {
			return errormessage.ErrBadRequest
		}
	case constant.BUYXPAYY:
		if p.MinQty == 0 || p.PayOnly == 0 {
			return errormessage.ErrBadRequest
		}
	case constant.DISCOUNT:
		if p.MinQty == 0 || p.Discount == 0 {
			return errormessage.ErrBadRequest
		}
	default:
		return errormessage.ErrBadRequest
	}

	return nil
}