package inventory

import (
	"OrderApp/persistency/table"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	salePrice := "0"
	mockProductRepository.On("GetProduct", "123").Return(&table.Product{ID: "123", Price: "-1", SalePrice: &salePrice}, nil)
	service := NewInventoryService(mockProductRepository)
	product, e := service.GetProduct("123")
	assert.NoError(t, e)
	assert.Equal(t, "123", product.ID)
	assert.NotEqual(t, "-1", product.GetFinalPrice())
}
