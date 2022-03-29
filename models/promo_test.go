package models

import "testing"

func TestPromo_validate(t *testing.T) {
	type fields struct {
		ID        int
		Sku       string
		PromoType string
		BonusSku  string
		MinQty    int
		PayOnly   int
		Discount  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Success",
			fields:  fields{
				ID:        0,
				Sku:       "JTX735",
				PromoType: "DISCOUNT",
				BonusSku:  "",
				MinQty:    5,
				PayOnly:   0,
				Discount:  10,
			},
			wantErr: false,
		},
		{
			name:    "Test Invalid promo type",
			fields:  fields{
				ID:        0,
				Sku:       "JTX735",
				PromoType: "DISCOUNT50",
				BonusSku:  "",
				MinQty:    5,
				PayOnly:   0,
				Discount:  10,
			},
			wantErr: true,
		},
		{
			name:    "Test invalid detail",
			fields:  fields{
				ID:        0,
				Sku:       "JTX735",
				PromoType: "BUY1GET1",
				BonusSku:  "",
				MinQty:    5,
				PayOnly:   0,
				Discount:  10,
			},
			wantErr: true,
		},
		{
			name:    "Test invalid detail",
			fields:  fields{
				ID:        0,
				Sku:       "JTX735",
				PromoType: "BUYXPAYY",
				BonusSku:  "TRJ287",
				MinQty:    5,
				PayOnly:   0,
				Discount:  10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Promo{
				ID:        tt.fields.ID,
				Sku:       tt.fields.Sku,
				PromoType: tt.fields.PromoType,
				BonusSku:  tt.fields.BonusSku,
				MinQty:    tt.fields.MinQty,
				PayOnly:   tt.fields.PayOnly,
				Discount:  tt.fields.Discount,
			}
			if err := p.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Promo.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
