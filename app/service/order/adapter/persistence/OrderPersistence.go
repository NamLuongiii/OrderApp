package persistence

import (
	"OrderApp/common/postgresql/table"
	model2 "OrderApp/service/order/application/domain/model"

	"gorm.io/gorm"
)

type OrderPersistence struct {
	db *gorm.DB
}

func NewOrderPersistence(db *gorm.DB) *OrderPersistence {
	return &OrderPersistence{db: db}
}

func (p *OrderPersistence) SaveOrder(order model2.Order) (string, error) {
	persistenceOrder := table.Order{
		Name:    order.GetName(),
		Total:   order.GetTotal().String(),
		Email:   order.GetEmail(),
		Phone:   order.GetPhone(),
		Address: order.GetAddress(),
		Note:    order.GetNote(),
		Status:  string(order.GetStatus()),
	}
	p.db.Create(&persistenceOrder)
	return persistenceOrder.ID, nil
}

func (p *OrderPersistence) GetOrder(id string) (model2.Order, error) {
	var persistenceOrder table.Order
	e := p.db.Where("id = ?", id).Preload("LineItems").First(&persistenceOrder).Error
	if e != nil {
		return nil, e
	}

	order, e := MapOrder(persistenceOrder)
	if e != nil {
		return nil, e
	}

	return order, nil
}

func (p *OrderPersistence) GetPaginatedOrders(page, limit int) ([]model2.Order, error) {
	var persistenceOrders []table.Order
	e := p.db.Limit(limit).
		Offset((page - 1) * limit).
		Preload("LineItems").
		Find(&persistenceOrders).Error
	if e != nil {
		return nil, e
	}

	var orders []model2.Order
	for _, persistenceOrder := range persistenceOrders {
		order, e := MapOrder(persistenceOrder)
		if e != nil {
			return nil, e
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (p *OrderPersistence) UpdateOrderStatus(orderId string, status model2.Status) error {
	err := p.db.
		Model(&table.Order{}).
		Where("id = ?", orderId).
		Update("status", string(status)).
		Error
	return err
}
