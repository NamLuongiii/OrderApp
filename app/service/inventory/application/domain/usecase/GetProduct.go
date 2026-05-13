package usecase

import (
	"OrderApp/service/inventory/application/domain/model"
	"OrderApp/service/inventory/application/port/out"
)

type GetProduct struct {
	persistenceProductPort out.PersistenceProductPort
}

func NewGetProduct(persistenceProductPort out.PersistenceProductPort) *GetProduct {
	return &GetProduct{
		persistenceProductPort: persistenceProductPort,
	}
}

func (g *GetProduct) GetProduct(id string) (*model.Product, error) {
	return g.persistenceProductPort.GetProduct(id)
}

func (g *GetProduct) GetPaginatedProducts(page, size int) ([]*model.Product, error) {
	products, err := g.persistenceProductPort.GetPaginatedProducts(page, size)
	if err != nil {
		return nil, err
	}
	return products, nil
}
