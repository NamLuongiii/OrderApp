package web

import (
	"OrderApp/order/application/port/in"

	"github.com/gin-gonic/gin"
)

type MarkOrderCompletedController struct {
	markOrderCompletedPort in.MarkOrderCompletedPort
}

func NewMarkOrderCompletedController(markOrderCompletedPort in.MarkOrderCompletedPort) *MarkOrderCompletedController {
	return &MarkOrderCompletedController{markOrderCompletedPort: markOrderCompletedPort}
}

func (c *MarkOrderCompletedController) BindHttpCall(r *gin.Engine) {
	r.PUT("/order/:orderId/completed", c.MarkOrderCompleted)
}

func (c *MarkOrderCompletedController) MarkOrderCompleted(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	err := c.markOrderCompletedPort.MarkOrderCompleted(orderId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to mark order completed"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Order completed successfully"})
}
