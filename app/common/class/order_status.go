package class

import (
	"OrderApp/common/msg"
	"errors"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusConfirmed OrderStatus = "CONFIRMED"
	StatusDelivered OrderStatus = "DELIVERED"
	StatusCanceled  OrderStatus = "CANCELED"
	StatusCompleted OrderStatus = "COMPLETED"
)

func ValidateOrderStatus(status string) error {
	switch OrderStatus(status) {
	case StatusPending,
		StatusConfirmed,
		StatusDelivered,
		StatusCanceled,
		StatusCompleted:
		return nil
	default:
		return errors.New(msg.InvalidOrderStatus)
	}
}
