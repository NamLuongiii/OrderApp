package inventory

import (
	"OrderApp/persistency/table"
)

func (s *ServiceImpl) CreateProduct(command CreateProductCommand) (string, error) {

	product := table.Product{
		Name:      command.Name,
		Price:     command.Price,
		SalePrice: &command.SalePrice,
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
	Price     int64
	SalePrice int64
}
