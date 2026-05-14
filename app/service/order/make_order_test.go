package order

import (
	"OrderApp/persistency/table"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeOrder_OverflowProof(t *testing.T) {
	// 1. Setup Mock
	mockOrderPersistency := new(MockOrderPersistency)
	mockLineItemPersistency := new(MockLineItemPersistency)
	mockProductPersistency := new(MockProductPersistency)

	// Biến để hứng giá trị Total thực tế mà Service tính toán được
	var actualTotalInRepo int64

	// Mock SaveOrder: Sử dụng MatchedBy để capture lại giá trị Total bị sai
	mockOrderPersistency.On("SaveOrder", mock.MatchedBy(func(o table.Order) bool {
		actualTotalInRepo = o.Total
		return o.Name == "name_test"
	})).Return("order_123", nil)

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
	service := NewService(mockOrderPersistency, mockLineItemPersistency, mockProductPersistency)
	orderId, e := service.MakeOrder(command)

	// 4. Assertions
	assert.NoError(t, e)
	assert.Equal(t, "order_123", orderId)

	// CHỈ RA TOTAL ORDER KHÔNG ĐÚNG:
	// Giá trị kỳ vọng toán học: (100 tỷ * 100 tỷ) + (200 tỷ * 100 tỷ) = 30,000,000,000,000,000,000,000
	// Con số này vượt xa MaxInt64 (9,223,372,036,854,775,807)

	// Khẳng định giá trị nhận được bị sai lệch (khác với 0 và khác với một con số nhỏ hợp lý)
	assert.Equal(t, int64(5594136148269072384), actualTotalInRepo, "Total bị overflow dẫn đến con số sai lệch")

	fmt.Printf("--- Kết quả Test ---\n")
	fmt.Printf("Total nhận được trong Repo: %v\n", actualTotalInRepo)
	fmt.Printf("--------------------\n")

	mockOrderPersistency.AssertExpectations(t)
}
