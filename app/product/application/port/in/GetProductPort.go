package in

import "OrderApp/product/application/domain/model"

type GetProductPort interface {
	GetPaginatedProducts(page, size int) ([]*model.Product, error)
	GetProduct(id string) (*model.Product, error)
}
