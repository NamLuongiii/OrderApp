package out

import "OrderApp/product/application/domain/model"

type PersistenceProductPort interface {
	SaveProduct(product *model.Product) error
	GetProduct(id string) (*model.Product, error)
	GetPaginatedProducts(page, size int) ([]*model.Product, error)
}
