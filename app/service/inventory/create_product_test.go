package inventory

import (
	"OrderApp/common/class"
	"OrderApp/persistency/table"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	price, e := class.NewMoney("-1")
	if e != nil {
		panic(e)
		return
	}

	salePrice, e := class.NewMoney("0")
	if e != nil {
		panic(e)
		return
	}
	command := CreateProductCommand{
		Name:      "test",
		Price:     price,
		SalePrice: &salePrice,
	}

	var salePriceStr string
	if command.SalePrice != nil {
		salePriceStr = (*command.SalePrice).String()
	}

	product := table.Product{
		Name:      command.Name,
		Price:     command.Price.String(),
		SalePrice: &salePriceStr,
	}
	mockProductRepository.On("SaveProduct", product).Return("1", nil)

	service := NewInventoryService(mockProductRepository)
	id, e := service.CreateProduct(command)

	assert.NoError(t, e)
	assert.Equal(t, "1", id)
}
