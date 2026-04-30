package out

import "OrderApp/order/application/domain/model"

type OrderPersistencePort interface {
	SaveOrder(order model.Order) (string, error)
	GetOrder(id string) (model.Order, error)
	GetPaginatedOrders(page, size int) ([]model.Order, error)
	UpdateOrderStatus(id string, status model.Status) error
}
