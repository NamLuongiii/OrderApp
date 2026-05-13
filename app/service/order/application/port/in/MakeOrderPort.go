package in

import (
	"OrderApp/service/order/application/domain/usecase"
)

type MakeOrderPort interface {
	MakeOrder(makeOrderCommand usecase.MakeOrderCommand) error
}
