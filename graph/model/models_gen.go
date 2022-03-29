// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cart struct {
	ID       int    `json:"ID"`
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type Carts struct {
	Contents []*NewCart `json:"contents"`
}

type NewCart struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type NewProduct struct {
	Sku      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type NewPromo struct {
	Sku         string       `json:"sku"`
	Promotype   string       `json:"promotype"`
	Promodetail *PromoDetail `json:"promodetail"`
}

type Product struct {
	ID       int     `json:"ID"`
	Sku      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Promo struct {
	ID        int    `json:"ID"`
	Promotype string `json:"promotype"`
	Bonussku  string `json:"bonussku"`
	Minqty    int    `json:"minqty"`
	Payonly   int    `json:"payonly"`
	Discount  int    `json:"discount"`
}

type PromoDetail struct {
	Bonussku *string `json:"bonussku"`
	Minqty   *int    `json:"minqty"`
	Payonly  *int    `json:"payonly"`
	Discount *int    `json:"discount"`
}

type ResponseCheckout struct {
	Total float64 `json:"total"`
}

type ResponseCreated struct {
	ID int `json:"ID"`
}

type ResponseUpdated struct {
	Status string `json:"status"`
}