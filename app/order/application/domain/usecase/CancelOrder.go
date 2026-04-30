package usecase

import (
	"OrderApp/order/application/domain/model"
	"OrderApp/order/application/port/out"
	"errors"
)

type CancelOrderUseCase struct {
	orderPersistencePort out.OrderPersistencePort
}

func NewCancelOrder(orderPersistencePort out.OrderPersistencePort) *CancelOrderUseCase {
	return &CancelOrderUseCase{orderPersistencePort: orderPersistencePort}
}

func (c *CancelOrderUseCase) CancelOrder(orderId string) error {
	order, e := c.orderPersistencePort.GetOrder(orderId)
	if e != nil {
		return e
	}

	if order.GetStatus() == model.COMPLETED || order.GetStatus() == model.CANCELLED {
		return errors.New("order has completed")
	}

	e = c.orderPersistencePort.UpdateOrderStatus(orderId, model.CANCELLED)
	return e
}
