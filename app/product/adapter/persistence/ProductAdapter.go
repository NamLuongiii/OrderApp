package persistence

import (
	"OrderApp/common/class"
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
	m1, _ := class.NewPositiveMoney("1000")
	m2, _ := class.NewPositiveMoney("800")
	return []*model.Product{
		model.NewProduct(
			"1",
			"p1",
			m1,
			&m2),
		model.NewProduct("2", "p2", m2, nil),
	}, nil
}
