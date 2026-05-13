package in

import "github.com/gin-gonic/gin"

type CheckRolePort interface {
	CheckRole() gin.HandlerFunc
}
