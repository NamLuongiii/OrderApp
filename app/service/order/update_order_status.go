package order

import (
	"OrderApp/common/class"
	"errors"
)

func (s ServiceImpl) UpdateOrderStatus(id string, status class.OrderStatus) error {
	switch status {
	case class.StatusConfirmed:
		return s.orderPersistency.ConfirmOrder(id)
	case class.StatusCanceled:
		return s.orderPersistency.CancelOrder(id)
	case class.StatusDelivered:
		return s.orderPersistency.MarkOrderAsDelivered(id)
	case class.StatusCompleted:
		return s.orderPersistency.MarkOrderAsCompleted(id)
	}

	return errors.New("invalid order status")
}
