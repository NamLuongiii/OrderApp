package persistence

import (
	"OrderApp/common/postgresql/table"
	"OrderApp/order/application/domain/model"

	"gorm.io/gorm"
)

type LineItemPersistence struct {
	db *gorm.DB
}

func NewLineItemPersistence(db *gorm.DB) *LineItemPersistence {
	return &LineItemPersistence{db: db}
}

func (p *LineItemPersistence) SaveLineItems(orderId string, lineItems []model.LineItem) error {
	var persistenceLineItem []table.LineItem
	for _, lineItem := range lineItems {
		persistenceLineItem = append(persistenceLineItem, table.LineItem{
			OrderID:     orderId,
			ProductID:   lineItem.GetProductID(),
			Quantity:    lineItem.GetProductQuantity(),
			Price:       lineItem.GetProductPrice().String(),
			Total:       lineItem.GetProductTotal().String(),
			ProductName: lineItem.GetProductName(),
		})
	}
	return p.db.CreateInBatches(&persistenceLineItem, len(persistenceLineItem)).Error // Insert 100 records mỗi batch
}
