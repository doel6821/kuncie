package models

import "github.com/doel6821/kuncie/package/errormessage"

// Product is represent table products
type Product struct {
	ID       int     `json:"id"`
	Sku      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	
}

// TableName products table name 
func (t *Product) TableName() string {
	return "products"
}

// NewProduct generate new product
func NewProduct(sku, name string , price float64, qtty int) (*Product, error) {
	product := &Product{
		ID:    0,
		Sku:   sku,
		Name:  name,
		Price: price,
		Quantity: qtty,
	}
	err := product.validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

// validate create product validation
func (p *Product) validate() error{
	if p.Name == "" || p.Sku == "" || p.Price == 0.0  || p.Quantity == 0{
		return errormessage.ErrBadRequest
	}
	return nil
}