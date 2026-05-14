package order

import (
	"OrderApp/common/class"
	"OrderApp/common/obj"
	"OrderApp/persistency"
	"OrderApp/persistency/table"
	"OrderApp/service/mail"
)

type Service interface {
	MakeOrder(command MakeOrderCommand) (string, error)
	GetOrder(id string) (*table.Order, error)
	GetOrderPagination(page, size int) ([]*table.Order, *obj.Pagination, error)
	UpdateOrderStatus(id string, status class.OrderStatus) error
}

type ServiceImpl struct {
	orderPersistency    persistency.OrderPersistency
	lineItemPersistency persistency.LineItemPersistency
	productPersistency  persistency.ProductPersistency
	mailService         mail.Service
}

func NewService(orderPersistency persistency.OrderPersistency,
	lineItemPersistency persistency.LineItemPersistency,
	productPersistency persistency.ProductPersistency,
	mailService mail.Service) Service {
	return &ServiceImpl{
		orderPersistency:    orderPersistency,
		lineItemPersistency: lineItemPersistency,
		productPersistency:  productPersistency,
		mailService:         mailService,
	}
}
