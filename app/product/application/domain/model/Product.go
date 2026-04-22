package model

import "OrderApp/common"

type Product struct {
	id   string
	name string

	price     *common.Money
	salePrice *common.Money
}

func NewProduct(id string, name string, price common.Money, salePrice *common.Money) *Product {
	return &Product{
		id:        id,
		name:      name,
		price:     &price,
		salePrice: salePrice,
	}
}

func (p Product) GetFinalPrice() common.Money {
	if p.salePrice != nil {
		return *p.salePrice
	}
	return *p.price
}
