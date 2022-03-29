package service

import (
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/errormessage"
	repo "github.com/doel6821/kuncie/repository"
)

type ProductServiceImpl struct {
	repo repo.ProductRepositoryInterface
}

func NewProductServiceImpl( repo repo.ProductRepositoryInterface) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo,
	}
}

// CreateProduct service for create new product
func (p *ProductServiceImpl) CreateProduct(req *model.NewProduct) (*models.Product, error) {
	newProduct , err := models.NewProduct(req.Sku, req.Name, float64(req.Price), req.Quantity)
	if err != nil {
		return nil, errormessage.ErrBadRequest
	}

	_ , err = p.repo.FindProductBySku(req.Sku)
	if err == nil {
		return nil, errormessage.ErrProductExist
	}

	product, err := p.repo.CreateProduct(newProduct)
	if err != nil {
		return nil, errormessage.ErrSaveDb
	}

	return product, nil

}

// UpdateProduct service for update product price or product quantity by sku
func (p *ProductServiceImpl)UpdateProduct(req *model.NewProduct) (*models.Product, error) {
	product , err := p.repo.FindProductBySku(req.Sku)
	if err != nil {
		return nil, errormessage.ErrNotFound
	}
	
	product.Price = req.Price
	product.Quantity = req.Quantity

	product, err = p.repo.CreateProduct(product)
	if err != nil {
		return nil, errormessage.ErrSaveDb
	}

	return product, nil

}

// FindProductBySku service for find product by sku
func (p *ProductServiceImpl) FindProductBySku(sku string) (*models.Product, error) {
	product , err := p.repo.FindProductBySku(sku)
	if err != nil {
		return nil, errormessage.ErrNotFound
	}

	return product, nil
}

// FindAllProduct find all product
func (p *ProductServiceImpl) FindAllProduct() ([]models.Product, error) {
	products , err := p.repo.FindAllProduct()
	if err != nil {
		return nil, errormessage.ErrNotFound
	}

	return products, nil
}