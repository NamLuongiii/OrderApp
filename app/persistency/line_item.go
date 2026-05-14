package persistency

import (
	"OrderApp/persistency/table"

	"gorm.io/gorm"
)

type LineItemPersistency interface {
	SaveLineItems(lineItems []*table.LineItem) error
}

type LineItemPersistenceImpl struct {
	db *gorm.DB
}

func NewLineItemPersistence(db *gorm.DB) *LineItemPersistenceImpl {
	return &LineItemPersistenceImpl{db: db}
}

func (p *LineItemPersistenceImpl) SaveLineItems(lineItems []*table.LineItem) error {
	return p.db.CreateInBatches(&lineItems, len(lineItems)).Error
}
