package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pass := "xxxxx"
	hash, err := hashPassword(pass)

	assert.NoError(t, err)
	assert.NotEqual(t, pass, hash)

	v := validatePassword(pass, hash)
	assert.True(t, v)
}
