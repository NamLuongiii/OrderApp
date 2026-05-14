package order

import (
	"OrderApp/common/msg"
	"OrderApp/persistency/table"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeOrder_OverflowProof(t *testing.T) {
	// 1. Setup Mock
	mockOrderPersistency := new(MockOrderPersistency)
	mockLineItemPersistency := new(MockLineItemPersistency)
	mockProductPersistency := new(MockProductPersistency)
	mockMailService := new(MockMailService)

	// Mock SaveLineItems: Chấp nhận bất kỳ mảng LineItems nào
	mockLineItemPersistency.On("SaveLineItems", mock.Anything).Return(nil)

	// Mock GetProductsByIDs: Trả về sản phẩm với giá cực lớn
	productIds := []string{"123", "456"}
	products := []*table.Product{
		{ID: "123", Price: 100000000000, Name: "Product 1"}, // 100 tỷ
		{ID: "456", Price: 200000000000, Name: "Product 2"}, // 200 tỷ
	}
	mockProductPersistency.On("GetProductsByIDs", productIds).Return(products, nil)

	// 2. Chuẩn bị Command với số lượng cực lớn để gây ra Overflow
	command := MakeOrderCommand{
		Products: []struct {
			ID       string
			Quantity int64
		}{
			{"123", 100000000000}, // 100 tỷ
			{"456", 100000000000}, // 100 tỷ
		},
		Customer: struct {
			Name    string
			Email   string
			Phone   string
			Address string
			Note    string
		}{
			Name:    "name_test",
			Email:   "a@gmail.com",
			Phone:   "0987654321",
			Address: "address_test",
			Note:    "note_test",
		},
	}

	// 3. Thực thi Service
	service := NewService(
		mockOrderPersistency,
		mockLineItemPersistency,
		mockProductPersistency,
		mockMailService,
	)
	_, e := service.MakeOrder(command)

	// 4. Assertions
	assert.Equal(t, e.Error(), msg.PriceValueTooLarge)

	mockOrderPersistency.AssertExpectations(t)
}
