package inventory

import (
	"OrderApp/common/class"
	"OrderApp/persistency/table"
)

func (s *ServiceImpl) CreateProduct(command CreateProductCommand) (string, error) {
	var salePriceStr string
	if command.SalePrice != nil {
		salePriceStr = (*command.SalePrice).String()
	}

	product := table.Product{
		Name:      command.Name,
		Price:     command.Price.String(),
		SalePrice: &salePriceStr,
	}

	id, e := s.productPersistency.SaveProduct(product)
	if e != nil {
		return "", e
	}
	if id != "" {
		return id, nil
	}

	return id, nil
}

type CreateProductCommand struct {
	Name      string
	Price     class.Money
	SalePrice *class.Money
}
