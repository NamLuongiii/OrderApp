package persistency

import (
	"OrderApp/common/class"
	"OrderApp/common/obj"
	"OrderApp/persistency/table"

	"gorm.io/gorm"
)

type OrderPersistency interface {
	SaveOrder(order table.Order) (string, error)
	GetOrder(id string) (*table.Order, error)
	GetPaginatedOrders(page, limit int) ([]*table.Order, *obj.Pagination, error)
	UpdateOrderStatus(orderId string, status string) error
	ConfirmOrder(orderId string) error
	CancelOrder(orderId string) error
	MarkOrderAsDelivered(orderId string) error
	MarkOrderAsCompleted(orderId string) error
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

func (p *OrderPersistenceImpl) GetPaginatedOrders(page, limit int) ([]*table.Order, *obj.Pagination, error) {
	var orders []*table.Order
	e := p.db.Limit(limit).
		Offset((page - 1) * limit).
		Preload("LineItems").
		Find(&orders).Error
	if e != nil {
		return nil, nil, e
	}

	var count int64
	e = p.db.Model(&table.Order{}).Count(&count).Error
	if e != nil {
		return nil, nil, e
	}
	pageNums := int(count / int64(limit))

	return orders, &obj.Pagination{
		Page:     page,
		PageNums: pageNums,
		PageSize: limit,
	}, nil
}

func (p *OrderPersistenceImpl) UpdateOrderStatus(orderId string, status string) error {
	err := p.db.
		Model(&table.Order{}).
		Where("id = ?", orderId).
		Update("status", status).
		Error
	return err
}

func (p *OrderPersistenceImpl) ConfirmOrder(orderId string) error {
	return p.db.
		Model(&table.Order{}).
		Where(&table.Order{ID: orderId, Status: string(class.StatusPending)}).
		Update("status", class.StatusConfirmed).
		Error
}

func (p *OrderPersistenceImpl) CancelOrder(orderId string) error {
	return p.db.
		Model(&table.Order{}).
		Where("id = ?", orderId).
		Where("status <> ?", string(class.StatusCompleted)).
		Where("status <> ?", string(class.StatusCompleted)).
		Where("status <> ?", string(class.StatusDelivered)).
		Update("status", class.StatusCanceled).
		Error
}

func (p *OrderPersistenceImpl) MarkOrderAsDelivered(orderId string) error {
	return p.db.
		Model(&table.Order{}).
		Where("id = ?", orderId).
		Where("status = ?", string(class.StatusConfirmed)).
		Update("status", class.StatusDelivered).
		Error
}

func (p *OrderPersistenceImpl) MarkOrderAsCompleted(orderId string) error {
	return p.db.
		Model(&table.Order{}).
		Where("id = ?", orderId).
		Where("status = ?", string(class.StatusDelivered)).
		Update("status", class.StatusCompleted).
		Error
}
