package persistency

import (
	"OrderApp/common/postgresql/table"

	"gorm.io/gorm"
)

type OrderPersistency interface {
	SaveOrder(order table.Order) (string, error)
	GetOrder(id string) (*table.Order, error)
	GetPaginatedOrders(page, limit int) ([]*table.Order, error)
	UpdateOrderStatus(orderId string, status string) error
}

type OrderPersistenceImpl struct {
	db *gorm.DB
}

func NewOrderPersistence(db *gorm.DB) OrderPersistency {
	return &OrderPersistenceImpl{db: db}
}

func (p *OrderPersistenceImpl) SaveOrder(order table.Order) (string, error) {
	p.db.Create(&order)
	return order.ID, nil
}

func (p *OrderPersistenceImpl) GetOrder(id string) (*table.Order, error) {
	var order table.Order
	e := p.db.
		Where("id = ?", id).
		Preload("LineItems").
		First(&order).
		Error
	if e != nil {
		return nil, e
	}
	return &order, nil
}

func (p *OrderPersistenceImpl) GetPaginatedOrders(page, limit int) ([]*table.Order, error) {
	var orders []*table.Order
	e := p.db.Limit(limit).
		Offset((page - 1) * limit).
		Preload("LineItems").
		Find(&orders).Error
	if e != nil {
		return nil, e
	}
	return orders, nil
}

func (p *OrderPersistenceImpl) UpdateOrderStatus(orderId string, status string) error {
	err := p.db.
		Model(&table.Order{}).
		Where("id = ?", orderId).
		Update("status", status).
		Error
	return err
}
