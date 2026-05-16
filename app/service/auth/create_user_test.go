package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	mockUser := new(MockUserPersistency)
	expectedError := errors.New("email is required")
	mockUser.On("CreateUser", mock.Anything).Return(expectedError)

	service := NewService(mockUser)

	command := CreateUserCommand{
		Name:     "test",
		Role:     "admin",
		Email:    "",
		Password: "password",
	}
	err := service.CreateUser(command)

	assert.Equal(t, expectedError, err)
}
