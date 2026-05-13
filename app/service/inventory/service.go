package inventory

import (
	"OrderApp/persistency"
	"OrderApp/persistency/table"
)

type Service interface {
	CreateProduct(command CreateProductCommand) (string, error)
	GetProduct(id string) (*table.Product, error)
	GetProductPagination(ids []string) (products []*table.Product, page int, size int, pageNum int, error error)
}

type ServiceImpl struct {
	productPersistency persistency.ProductPersistency
}

func NewInventoryService(productPersistency persistency.ProductPersistency) Service {
	return ServiceImpl{
		productPersistency: productPersistency,
	}
}
