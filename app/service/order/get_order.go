package order

import "OrderApp/persistency/table"

func (s ServiceImpl) GetOrder(id string) (*table.Order, error) {
	return s.orderPersistency.GetOrder(id)
}
