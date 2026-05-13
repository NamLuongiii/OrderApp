package web

import (
	"OrderApp/service/auth/application/port/in"

	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	getUserPort in.GetUserPort
}

func NewGetUserController(getUserPort in.GetUserPort) *GetUserController {
	return &GetUserController{
		getUserPort: getUserPort,
	}
}

func (c *GetUserController) BindHttpCall(r *gin.Engine) {
	r.GET("/user/:id", c.GetUser)
}

func (c *GetUserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, e := c.getUserPort.GetUser(id)

	if e != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(200, user)
}
