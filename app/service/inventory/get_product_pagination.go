package inventory

import (
	"OrderApp/persistency/table"
)

func (s ServiceImpl) GetProductPagination(ids []string) (
	products []*table.Product, page int, size int, pageNum int, error error) {
	return nil, 0, 0, 0, nil
}
