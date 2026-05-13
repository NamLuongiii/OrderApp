package web

import (
	"OrderApp/service/order/application/port/in"

	"github.com/gin-gonic/gin"
)

type MarkOrderDeliveredController struct {
	markOrderDeliveredPort in.MarkOrderDeliveredPort
}

func NewMarkOrderDeliveredController(markOrderDeliveredPort in.MarkOrderDeliveredPort) *MarkOrderDeliveredController {
	return &MarkOrderDeliveredController{markOrderDeliveredPort: markOrderDeliveredPort}
}

func (c *MarkOrderDeliveredController) BindHttpCall(r *gin.Engine) {
	r.PUT("/order/:orderId/delivered", c.MarkOrderDelivered)
}

func (c *MarkOrderDeliveredController) MarkOrderDelivered(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	err := c.markOrderDeliveredPort.MarkOrderDelivered(orderId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to mark order delivered"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Order delivered successfully"})
}
