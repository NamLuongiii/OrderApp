package in

import (
	"OrderApp/service/order/application/domain/model"
)

type GetOrderPort interface {
	GetOrder(id string) (model.Order, error)
	GetPaginatedOrders(page int, size int) ([]model.Order, error)
}
