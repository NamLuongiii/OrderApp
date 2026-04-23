package usecase

import (
	"OrderApp/product/application/domain/model"
	"OrderApp/product/application/port/out"
)

type GetProduct struct {
	persistenceProductPort *out.PersistenceProductPort
}

func NewGetProduct(persistenceProductPort *out.PersistenceProductPort) *GetProduct {
	return &GetProduct{
		persistenceProductPort: persistenceProductPort,
	}
}

func (g *GetProduct) GetProduct(id string) (*model.Product, error) {
	return nil, nil
}

func (g *GetProduct) GetPaginatedProducts(page, size int) ([]*model.Product, error) {
	return nil, nil
}
