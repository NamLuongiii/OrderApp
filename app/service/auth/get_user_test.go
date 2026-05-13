package auth

import (
	"OrderApp/persistency/table"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	mockUser := new(MockUserPersistency)
	expectedUser := &table.User{
		ID:    "123",
		Email: "test@example.com",
	}

	mockUser.On("GetUser", "123").Return(expectedUser, nil)

	service := NewService(mockUser)
	user, err := service.GetUser("123")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockUser.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	mockUser := new(MockUserPersistency)
	expectedError := errors.New("user not found")

	mockUser.On("GetUser", "999").Return(nil, expectedError)

	service := NewService(mockUser)
	user, err := service.GetUser("999")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, expectedError, err)
	mockUser.AssertExpectations(t)
}
