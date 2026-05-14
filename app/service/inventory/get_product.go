package inventory

import (
	"OrderApp/persistency/table"
)

func (s *ServiceImpl) GetProduct(id string) (*table.Product, error) {
	return s.productPersistency.GetProduct(id)
}
