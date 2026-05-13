package out

import (
	model2 "OrderApp/service/order/application/domain/model"
)

type OrderPersistencePort interface {
	SaveOrder(order model2.Order) (string, error)
	GetOrder(id string) (model2.Order, error)
	GetPaginatedOrders(page, size int) ([]model2.Order, error)
	UpdateOrderStatus(id string, status model2.Status) error
}
