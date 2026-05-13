package auth

import (
	"OrderApp/common/postgresql/table"

	"github.com/stretchr/testify/mock"
)

// MockUserPersistency is a mock implementation using testify/mock
type MockUserPersistency struct {
	mock.Mock
}

func (m *MockUserPersistency) CreateUser(user table.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserPersistency) GetUser(id string) (*table.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*table.User), args.Error(1)
}

func (m *MockUserPersistency) GetUserByEmail(email string) (*table.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*table.User), args.Error(1)
}
