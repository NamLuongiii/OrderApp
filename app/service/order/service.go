package order

import "OrderApp/persistency"

type Service interface {
	MakeOrder(command MakeOrderCommand) (string, error)
}

type ServiceImpl struct {
	orderPersistency    persistency.OrderPersistency
	lineItemPersistency persistency.LineItemPersistency
	productPersistency  persistency.ProductPersistency
}

func NewService(orderPersistency persistency.OrderPersistency,
	lineItemPersistency persistency.LineItemPersistency,
	productPersistency persistency.ProductPersistency) Service {
	return &ServiceImpl{
		orderPersistency:    orderPersistency,
		lineItemPersistency: lineItemPersistency,
		productPersistency:  productPersistency,
	}
}
