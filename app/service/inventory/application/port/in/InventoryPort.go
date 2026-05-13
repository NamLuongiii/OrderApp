package in

import (
	"OrderApp/service/inventory/application/domain/model"
)

type InventoryPort interface {
	GetProductsBatch(ids []string) ([]*model.Product, error)
}
