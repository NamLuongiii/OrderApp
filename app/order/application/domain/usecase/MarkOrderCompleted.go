package usecase

import (
	"OrderApp/order/application/domain/model"
	"OrderApp/order/application/port/out"
	"errors"
)

type MarkOrderCompletedUseCase struct {
	orderPersistencePort out.OrderPersistencePort
}

func NewMarkOrderCompletedUseCase(orderPersistencePort out.OrderPersistencePort) *MarkOrderCompletedUseCase {
	return &MarkOrderCompletedUseCase{
		orderPersistencePort: orderPersistencePort,
	}
}

func (c *MarkOrderCompletedUseCase) MarkOrderCompleted(orderId string) error {
	order, e := c.orderPersistencePort.GetOrder(orderId)
	if e != nil {
		return e
	}

	if order.GetStatus() != model.DELIVERING {
		return errors.New("order status must be DELIVERING")
	}

	e = c.orderPersistencePort.UpdateOrderStatus(orderId, model.COMPLETED)

	return e
}
