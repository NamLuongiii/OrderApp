package persistence

import (
	"OrderApp/product/application/domain/model"
	"fmt"
)

type ProductAdapter struct {
}

func ProductAdapterImpl() *ProductAdapter {
	return &ProductAdapter{}
}

func (p *ProductAdapter) SaveProduct(product *model.Product) error {
	fmt.Println("Save product ", product.GetName(), product.GetFinalPrice())
	return nil
}

func (p *ProductAdapter) GetProduct(id string) (*model.Product, error) {
	return nil, nil
}

func (p *ProductAdapter) GetPaginatedProducts(page, size int) ([]*model.Product, error) {
	return nil, nil
}
