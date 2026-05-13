package auth

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetUser(t *testing.T) {
	service := NewService()
	msg := service.getUser()

	assert.Equal(t, "hello worlds", msg)
}
