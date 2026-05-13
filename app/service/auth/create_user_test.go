package auth

import (
	"OrderApp/persistency/table"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockUser := new(MockUserPersistency)

	user := table.User{
		Name:     "test",
		Role:     "admin",
		Email:    "",
		Password: "",
	}

	expectedError := errors.New("email is required")
	mockUser.On("CreateUser", user).Return(expectedError)

	service := NewService(mockUser)

	command := CreateUserCommand{
		Name:     "test",
		Role:     "admin",
		Email:    "",
		Password: "",
	}
	err := service.CreateUser(command)

	assert.Equal(t, expectedError, err)
}
