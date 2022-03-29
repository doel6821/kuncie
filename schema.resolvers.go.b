package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "fmt"

	"github.com/doel6821/kuncie/graph/generated"
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/repository"
	"github.com/doel6821/kuncie/service"
)

var (
	productRepo     = repository.NewProductRepositoryImpl()
	productService  = service.NewProductServiceImpl(productRepo)
	promoRepo       = repository.NewPromoRepositoryImpl()
	promoService    = service.NewPromoServiceImpl(promoRepo)
	cartRepo        = repository.NewCartRepositoryImpl()
	cartService     = service.NewCartServiceImpl(cartRepo)
	checkoutService = service.NewCheckoutServiceImpl(productRepo, promoRepo, cartRepo)
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.ResponseCreated, error) {
	product, err := productService.CreateProduct(input)
	if err != nil {
		return nil, err
	}
	return &model.ResponseCreated{
		ID: product.ID,
	}, nil
}

func (r *mutationResolver) CreatePromo(ctx context.Context, input *model.NewPromo) (*model.ResponseCreated, error) {
	promo, err := promoService.CreatePromo(input)
	if err != nil {
		return nil, err
	}
	return &model.ResponseCreated{
		ID: promo.ID,
	}, nil
}

func (r *mutationResolver) AddCart(ctx context.Context, input *model.NewCart) (*model.Cart, error) {
	cart, err := cartService.CreateCart(input)
	if err != nil {
		return nil, err
	}
	return &model.Cart{
		ID:       cart.ID,
		Sku:      cart.Sku,
		Quantity: cart.Quantity,
	}, nil
}

func (r *queryResolver) Product(ctx context.Context, sku string) (*model.Product, error) {
	product, err := productService.FindProductBySku(sku)
	if err != nil {
		return nil, err
	}
	return &model.Product{
		ID:       product.ID,
		Sku:      product.Sku,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := productService.FindAllProduct()
	if err != nil {
		return nil, err
	}
	
	var res []*model.Product
	for _, p := range products {
		res = append(res, &model.Product{
			ID:       p.ID,
			Sku:      p.Sku,
			Name:     p.Name,
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}
	
	return res, nil
}

func (r *queryResolver) Promo(ctx context.Context, sku string) (*model.Promo, error) {
	promo, err := promoService.GetPromoBySku(sku)
	if err != nil {
		return nil, err
	}
	return &model.Promo{
		ID: promo.ID,
		Promotype: promo.PromoType,
		Bonussku:  promo.BonusSku,
		Minqty:    promo.MinQty,
		Payonly:   promo.PayOnly,
		Discount:  promo.Discount,
	}, nil
}

func (r *queryResolver) Promolist(ctx context.Context) ([]*model.Promo, error) {
	promolist, err := promoService.GetAllPromo()
	if err != nil {
		return nil, err
	}

	var res []*model.Promo

	for _, p := range promolist {
		res = append(res, &model.Promo{
			ID: p.ID,
			Promotype: p.PromoType,
			Bonussku:  p.BonusSku,
			Minqty:    p.MinQty,
			Payonly:   p.PayOnly,
			Discount:  p.Discount,
		})
	}
	return res, nil
}

func (r *queryResolver) Checkout(ctx context.Context, input *model.Carts) (*model.ResponseCheckout, error) {
	pay, err := checkoutService.Checkout(input)
	if err != nil {
		return &model.ResponseCheckout{}, err
	}

	return &model.ResponseCheckout{
		Total: pay,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
