package usecase

import (
	"OrderApp/inventory/application/domain/model"
	"OrderApp/inventory/application/port/out"
)

type InventoryPortImpl struct {
	persistenceProductPort out.PersistenceProductPort
}

func NewInventoryPort(persistenceProductPort out.PersistenceProductPort) *InventoryPortImpl {
	return &InventoryPortImpl{persistenceProductPort: persistenceProductPort}
}

func (service *InventoryPortImpl) GetProductsBatch(ids []string) ([]*model.Product, error) {
	return service.persistenceProductPort.GetProductsByIDs(ids)
}
