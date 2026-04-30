package in

import (
	"OrderApp/order/application/domain/usecase"
)

type MakeOrderPort interface {
	MakeOrder(makeOrderCommand usecase.MakeOrderCommand) error
}
