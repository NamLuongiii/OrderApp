package inventory

import (
	"OrderApp/common/obj"
	"OrderApp/persistency/table"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductPagination(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.On("GetPaginatedProducts", 1, 0).Return([]*table.Product{}, &obj.Pagination{}, errors.New("internal server error"))

	service := NewInventoryService(mockProductRepository)
	_, _, e := service.GetProductPagination(1, 0)
	assert.Error(t, e)
}
