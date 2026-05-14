package inventory

import (
	"OrderApp/common/obj"
	"OrderApp/persistency"
	"OrderApp/persistency/table"
)

type Service interface {
	CreateProduct(command CreateProductCommand) (string, error)
	GetProduct(id string) (*table.Product, error)
	GetProductPagination(page, size int) ([]*table.Product, *obj.Pagination, error)
}

type ServiceImpl struct {
	productPersistency persistency.ProductPersistency
}

func NewInventoryService(productPersistency persistency.ProductPersistency) Service {
	return &ServiceImpl{
		productPersistency: productPersistency,
	}
}
