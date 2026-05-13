package usecase

import (
	"OrderApp/service/auth/application/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckRoleMiddleware struct {
}

func NewCheckRoleMiddleware() *CheckRoleMiddleware {
	return &CheckRoleMiddleware{}
}

func (c *CheckRoleMiddleware) CheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims, err := model.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
