package in

import "OrderApp/inventory/application/domain/model"

type InventoryPort interface {
	GetProductsBatch(ids []string) ([]*model.Product, error)
}
