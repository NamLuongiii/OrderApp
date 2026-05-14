package order

import (
	"OrderApp/common/obj"
	"OrderApp/persistency/table"
)

func (s ServiceImpl) GetOrderPagination(page, size int) ([]*table.Order, *obj.Pagination, error) {
	return s.orderPersistency.GetPaginatedOrders(page, size)
}
