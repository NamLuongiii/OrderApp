package order

import (
	"OrderApp/common/obj"
	"OrderApp/persistency/table"

	"github.com/stretchr/testify/mock"
)

// ////////////////////////////////////////
//
//	Mock Order Persistency layer         //
//
// ////////////////////////////////////////
type MockOrderPersistency struct {
	mock.Mock
}

func (m *MockOrderPersistency) SaveOrder(order table.Order) (string, error) {
	args := m.Called(order)
	return args.String(0), args.Error(1)
}

func (m *MockOrderPersistency) GetOrder(id string) (*table.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*table.Order), args.Error(1)
}

func (m *MockOrderPersistency) GetPaginatedOrders(page, limit int) ([]*table.Order, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]*table.Order), args.Error(1)
}

func (m *MockOrderPersistency) UpdateOrderStatus(orderId string, status string) error {
	args := m.Called(orderId, status)
	return args.Error(0)
}

// ////////////////////////////////////////
//
//	Mock LineItem Persistency layer      //
//
// ////////////////////////////////////////
type MockLineItemPersistency struct {
	mock.Mock
}

func (m *MockLineItemPersistency) SaveLineItems(lineItems []*table.LineItem) error {
	args := m.Called(lineItems)
	return args.Error(0)
}

// ////////////////////////////////////////
//
//	Mock Product            //
//	Persistency layer        //
//	                //
//
// ////////////////////////////////////////
type MockProductPersistency struct {
	mock.Mock
}

func (m *MockProductPersistency) SaveProduct(product table.Product) (string, error) {
	args := m.Called(product)
	return args.String(0), args.Error(1)
}

func (m *MockProductPersistency) GetProduct(id string) (*table.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*table.Product), args.Error(1)
}

func (m *MockProductPersistency) GetPaginatedProducts(page, size int) ([]*table.Product, *obj.Pagination, error) {
	args := m.Called(page, size)
	return args.Get(0).([]*table.Product), args.Get(1).(*obj.Pagination), args.Error(2)
}

func (m *MockProductPersistency) GetProductsByIDs(ids []string) ([]*table.Product, error) {
	args := m.Called(ids)
	return args.Get(0).([]*table.Product), args.Error(1)
}
