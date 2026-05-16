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

	v := checkPasswordCorrect(pass, hash)
	assert.True(t, v)

	p2 := "12345"
	hash2 := "$2a$10$hZsyprqgAlc5eQFSGSnOT.knP73lgV2mUwKAWtdrq7vzZHCHRuk.u"
	v2 := checkPasswordCorrect(p2, hash2)
	assert.True(t, v2)
}
