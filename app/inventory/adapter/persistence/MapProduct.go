package persistence

import (
	"OrderApp/common/class"
	"OrderApp/common/postgresql/table"
	"OrderApp/inventory/application/domain/model"
)

func MapProduct(table table.Product) (*model.Product, error) {
	price, _ := class.NewPositiveMoney(table.Price)
	var salePrice *class.Money = nil

	if table.SalePrice != nil && *table.SalePrice != "" {
		sp, e := class.NewPositiveMoney(*table.SalePrice)
		if e != nil {
			return nil, e
		}
		salePrice = &sp
	}
	return model.NewProduct(
		table.ID,
		table.Name,
		price,
		salePrice), nil
}
