package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwtToken(t *testing.T) {
	claims := Token{
		UserID: "123",
		Role:   "admin",
	}

	token, e := generateJwtToken(claims.UserID, claims.Role)

	assert.Equal(t, e, nil)
	assert.NotEmpty(t, token)

	s, e := verifyToken(token)
	assert.Equal(t, e, nil)
	assert.Equal(t, claims.UserID, s.UserID)
	assert.Equal(t, claims.Role, s.Role)

	expected := time.Now().Add(24 * time.Hour).Truncate(time.Minute)
	actual := s.RegisteredClaims.ExpiresAt.Time.Truncate(time.Minute)
	assert.Equal(t, expected, actual)

}
