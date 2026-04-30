package out

import "OrderApp/inventory/application/domain/model"

type PersistenceProductPort interface {
	SaveProduct(product *model.Product) error
	GetProduct(id string) (*model.Product, error)
	GetPaginatedProducts(page, size int) ([]*model.Product, error)
	GetProductsByIDs(ids []string) ([]*model.Product, error)
}
