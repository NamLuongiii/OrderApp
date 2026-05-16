package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCheckRoleMiddleware_NoToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/test", CheckRole(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	c.Request = httptest.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCheckRoleMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/test", CheckRole(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "invalid-token")
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCheckRoleMiddleware_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	// Generate valid token
	token, err := generateJwtToken("user123", RoleAdmin)
	assert.NoError(t, err)

	r.GET("/test", CheckRole(), func(c *gin.Context) {
		userId, exists := c.Get("userId")
		assert.True(t, exists)
		assert.Equal(t, "user123", userId)

		role, exists := c.Get("role")
		assert.True(t, exists)
		assert.Equal(t, RoleAdmin, role)

		c.JSON(200, gin.H{"message": "success"})
	})

	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", token)
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, http.StatusOK, w.Code)
}
