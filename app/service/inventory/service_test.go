package inventory

import (
	"OrderApp/common/obj"
	"OrderApp/persistency/table"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) SaveProduct(product table.Product) (string, error) {
	args := m.Called(product)
	return args.String(0), args.Error(1)
}

func (m *MockProductRepository) GetProduct(id string) (*table.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*table.Product), args.Error(1)
}

func (m *MockProductRepository) GetPaginatedProducts(page, size int) ([]*table.Product, *obj.Pagination, error) {
	args := m.Called(page, size)
	return args.Get(0).([]*table.Product), args.Get(1).(*obj.Pagination), args.Error(2)
}

func (m *MockProductRepository) GetProductsByIDs(ids []string) ([]*table.Product, error) {
	args := m.Called(ids)
	return args.Get(0).([]*table.Product), args.Error(1)
}
