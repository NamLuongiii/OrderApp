package usecase

import (
	"OrderApp/order/application/domain/model"
	"OrderApp/order/application/port/out"
)

type GetOrder struct {
	orderPersistencePort out.OrderPersistencePort
}

func NewGetOrder(orderPersistencePort out.OrderPersistencePort) *GetOrder {
	return &GetOrder{orderPersistencePort: orderPersistencePort}
}

func (g *GetOrder) GetOrder(id string) (model.Order, error) {
	return g.orderPersistencePort.GetOrder(id)
}

func (g *GetOrder) GetPaginatedOrders(page, size int) ([]model.Order, error) {
	return g.orderPersistencePort.GetPaginatedOrders(page, size)
}
