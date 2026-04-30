package model

import (
	"OrderApp/common/class"
)

type Product struct {
	id   string
	name string

	price     *class.Money
	salePrice *class.Money
}

func NewProduct(id string, name string, price class.Money, salePrice *class.Money) *Product {
	return &Product{
		id:        id,
		name:      name,
		price:     &price,
		salePrice: salePrice,
	}
}

func NewProductWithoutId(name string, price class.Money, salePrice *class.Money) *Product {
	return &Product{
		name:      name,
		price:     &price,
		salePrice: salePrice,
	}
}

func (p *Product) GetId() string { return p.id }

func (p *Product) GetName() string { return p.name }

func (p *Product) GetPrice() class.Money { return *p.price }

func (p *Product) GetSalePrice() *class.Money { return p.salePrice }

func (p *Product) GetFinalPrice() class.Money {
	if p.salePrice != nil {
		return *p.salePrice
	}
	return *p.price
}
