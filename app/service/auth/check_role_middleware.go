package auth

import (
	"OrderApp/common/msg"
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
			c.JSON(http.StatusUnauthorized, msg.InvalidToken)
			c.Abort()
			return
		}

		claims, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msg.InvalidToken)
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
