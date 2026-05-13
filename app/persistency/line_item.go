package persistency

import (
	"OrderApp/common/postgresql/table"
	"OrderApp/service/order/application/domain/model"

	"gorm.io/gorm"
)

type LineItemPersistency interface {
	SaveLineItems(orderId string, lineItems []model.LineItem) error
}

type LineItemPersistenceImpl struct {
	db *gorm.DB
}

func NewLineItemPersistence(db *gorm.DB) *LineItemPersistenceImpl {
	return &LineItemPersistenceImpl{db: db}
}

func (p *LineItemPersistenceImpl) SaveLineItems(lineItems []table.LineItem) error {
	return p.db.CreateInBatches(&lineItems, len(lineItems)).Error
}
