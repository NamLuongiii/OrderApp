package usecase

import (
	"OrderApp/service/order/application/domain/model"
	"OrderApp/service/order/application/port/out"
	"errors"
)

type ConfirmOrderUseCase struct {
	orderPersistencePort out.OrderPersistencePort
}

func NewConfirmOrder(orderPersistencePort out.OrderPersistencePort) *ConfirmOrderUseCase {
	return &ConfirmOrderUseCase{orderPersistencePort: orderPersistencePort}
}

func (c *ConfirmOrderUseCase) ConfirmOrder(orderId string) error {
	order, err := c.orderPersistencePort.GetOrder(orderId)
	if err != nil {
		return err
	}

	if order.GetStatus() != model.PROCESSING {
		return errors.New("order status must be PROCESSING")
	}

	return c.orderPersistencePort.UpdateOrderStatus(orderId, model.CONFIRMED)
}
