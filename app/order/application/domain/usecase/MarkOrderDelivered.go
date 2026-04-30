package usecase

import (
	"OrderApp/order/application/domain/model"
	"OrderApp/order/application/port/out"
	"errors"
)

type MarkOrderDeliveredUseCase struct {
	orderPersistencePort out.OrderPersistencePort
}

func NewMarkOrderDeliveredUseCase(orderPersistencePort out.OrderPersistencePort) *MarkOrderDeliveredUseCase {
	return &MarkOrderDeliveredUseCase{orderPersistencePort: orderPersistencePort}

}

func (c *MarkOrderDeliveredUseCase) MarkOrderDelivered(orderId string) error {
	order, e := c.orderPersistencePort.GetOrder(orderId)
	if e != nil {
		return e
	}

	if order.GetStatus() != model.CONFIRMED {
		return errors.New("order status must be confirmed")
	}

	e = c.orderPersistencePort.UpdateOrderStatus(orderId, model.DELIVERING)

	return e

}
