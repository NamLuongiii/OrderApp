package inventory

import (
	"OrderApp/persistency/table"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockProductRepository := new(MockProductRepository)

	command := CreateProductCommand{
		Name:      "test",
		Price:     1,
		SalePrice: -1,
	}

	var salePrice int64 = -1
	product := table.Product{
		Name:      command.Name,
		Price:     1,
		SalePrice: &salePrice,
	}
	mockProductRepository.On("SaveProduct", product).Return("1", nil)

	service := NewInventoryService(mockProductRepository)
	id, e := service.CreateProduct(command)

	assert.NoError(t, e)
	assert.Equal(t, "1", id)
	assert.NotEqual(t, "-1", product.GetFinalPrice())
}
