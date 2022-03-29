package service

import (
	"github.com/doel6821/kuncie/graph/model"
	"github.com/doel6821/kuncie/models"
	"github.com/doel6821/kuncie/package/errormessage"
	repo "github.com/doel6821/kuncie/repository"
)

// PromoServiceImpl promo service implementation
type PromoServiceImpl struct {
	repo repo.PromoRepositoryInterface
}

// NewPromoServiceImpl create new promo implementation
func NewPromoServiceImpl(repo repo.PromoRepositoryInterface) *PromoServiceImpl {
	return &PromoServiceImpl{
		repo,
	}
}

// CreatePromo create new promo
func (p *PromoServiceImpl) CreatePromo(req *model.NewPromo) (*models.Promo, error) {
	_ , err := p.repo.FindPromoBYSku(req.Sku)
	if err == nil {
		return nil, errormessage.ErrProductExist
	}


	promo , err := models.NewPromo(req.Sku, req.Promotype, req.Promodetail)
	if err != nil {
		return nil, err
	}


	promo, err = p.repo.CreatePromo(promo)
	if err != nil {
		return nil, errormessage.ErrSaveDb
	}

	return promo, nil
}

// UpdatePromo update promo
func (p *PromoServiceImpl) UpdatePromo(req *model.NewPromo) (*models.Promo, error) {
	_ , err := models.NewPromo(req.Sku, req.Promotype, req.Promodetail)
	if err != nil {
		return nil, err
	}

	promo , err := p.repo.FindPromoBYSku(req.Sku)
	if err != nil {
		return nil, errormessage.ErrNotFound
	}

	promo.PromoType = req.Promotype
	promo.BonusSku = *req.Promodetail.Bonussku
	promo.Discount = *req.Promodetail.Discount
	promo.MinQty = *req.Promodetail.Minqty
	promo.PayOnly = *req.Promodetail.Payonly

	promo, err = p.repo.CreatePromo(promo)
	if err != nil {
		return nil, errormessage.ErrSaveDb
	}

	return promo, nil
}


// GetPromoBySku get promo by sku
func (p *PromoServiceImpl) GetPromoBySku( sku string) (*models.Promo, error) {
	promo , err := p.repo.FindPromoBYSku(sku)
	if err != nil {
		return nil, errormessage.ErrNotFound
	}

	return promo, nil
}

// GetAllPromo get all promo
func (p *PromoServiceImpl) GetAllPromo() ([]models.Promo, error) {
	promolist , err := p.repo.FindAllPromo()
	if err != nil {
		return nil, errormessage.ErrNotFound
	}

	return promolist, nil
}
