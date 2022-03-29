package models

import (
	"reflect"
	"testing"
)

func TestProduct_validate(t *testing.T) {
	type fields struct {
		ID       int
		Sku      string
		Name     string
		Price    float64
		Quantity int
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
				ID:       0,
				Sku:      "HVE123",
				Name:     "Laptop Case",
				Price:    50000,
				Quantity: 5,
			},
			wantErr: false,
		},
		{
			name: "Test Empty value",
			fields: fields{
				ID:       0,
				Sku:      "",
				Name:     "",
				Price:    0,
				Quantity: 0,
			},
			wantErr: true,
		},
		{
			name: "Test Empty quantity",
			fields: fields{
				ID:       0,
				Sku:      "H1E123",
				Name:     "Case HP",
				Price:    45000,
				Quantity: 0,
			},
			wantErr: true,
		},
		{
			name: "Test Empty price",
			fields: fields{
				ID:       0,
				Sku:      "H1E123",
				Name:     "Case HP Sungsam",
				Price:    0,
				Quantity: 10,
			},
			wantErr: true,
		},
		{
			name: "Test Empty quantity",
			fields: fields{
				ID:       0,
				Sku:      "R1E123",
				Name:     "Case HP Novia",
				Price:    45000,
				Quantity: 0,
			},
			wantErr: true,
		},
		{
			name: "Test Empty sku",
			fields: fields{
				ID:       0,
				Sku:      "",
				Name:     "Case HP LaGi",
				Price:    25000,
				Quantity: 3,
			},
			wantErr: true,
		},
		{
			name: "Test Empty name",
			fields: fields{
				ID:       0,
				Sku:      "R1E121",
				Name:     "",
				Price:    35000,
				Quantity: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:       tt.fields.ID,
				Sku:      tt.fields.Sku,
				Name:     tt.fields.Name,
				Price:    tt.fields.Price,
				Quantity: tt.fields.Quantity,
			}
			if err := p.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Product.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewProduct(t *testing.T) {
	type args struct {
		sku   string
		name  string
		price float64
		qtty  int
	}
	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Success",
			args:    args{
				sku:   "HVE123",
				name:  "Laptop Case",
				price: 50000,
				qtty:  5,
			},
			want:    &Product{
				ID:       0,
				Sku:      "HVE123",
				Name:     "Laptop Case",
				Price:    50000,
				Quantity: 5,
			},
			wantErr: false,
		},
		{
			name:    "Test Empty Parameter",
			args:    args{
				sku:   "",
				name:  "",
				price: 0,
				qtty:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test Empty sku",
			args:    args{
				sku:   "",
				name:  "Casing HP",
				price: 50000,
				qtty:  3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test Empty name",
			args:    args{
				sku:   "BVE123",
				name:  "",
				price: 50000,
				qtty:  8,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test Empty price",
			args:    args{
				sku:   "HVE123",
				name:  "Chanrger Universal",
				price: 0,
				qtty:  10,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test Empty quantity",
			args:    args{
				sku:   "IYE413",
				name:  "Kabel Type C",
				price: 30000,
				qtty:  0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.sku, tt.args.name, tt.args.price, tt.args.qtty)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
