package inventory

import (
	"OrderApp/common/obj"
	"OrderApp/persistency/table"
)

func (s *ServiceImpl) GetProductPagination(page, size int) ([]*table.Product, *obj.Pagination, error) {
	return s.productPersistency.GetPaginatedProducts(page, size)
}
