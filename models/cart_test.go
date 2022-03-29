package models

import (
	"reflect"
	"testing"
)

func TestCart_validate(t *testing.T) {
	type fields struct {
		ID       int
		Sku      string
		Quantity int
		Status   int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Success",
			fields: fields{
				Sku:      "PS123A",
				Quantity: 1,
			},
			wantErr: false,
		},
		{
			name:    "Test empty parameter",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "Test empty sku",
			fields: fields{
				Sku:      "",
				Quantity: 3,
			},
			wantErr: true,
		},
		{
			name: "Test empty qtty",
			fields: fields{
				Sku:      "SKU123",
				Quantity: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				ID:       tt.fields.ID,
				Sku:      tt.fields.Sku,
				Quantity: tt.fields.Quantity,
				Status:   tt.fields.Status,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Cart.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCart(t *testing.T) {
	type args struct {
		quantity int
		sku      string
	}
	tests := []struct {
		name    string
		args    args
		want    *Cart
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Success",
			args:    args{
				quantity: 2,
				sku:      "SKU123",
			},
			want:    &Cart{
				ID:       0,
				Sku:      "SKU123",
				Quantity: 2,
				Status:   0,
			},
			wantErr: false,
		},
		{
			name:    "Test Empty Parameter",
			args:    args{
				quantity: 0,
				sku:      "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test Empty quantity",
			args:    args{
				quantity: 0,
				sku:      "ABC123",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test Empty Sku",
			args:    args{
				quantity: 5,
				sku:      "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCart(tt.args.quantity, tt.args.sku)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
