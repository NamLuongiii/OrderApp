package web

import (
	"OrderApp/order/application/port/in"

	"github.com/gin-gonic/gin"
)

type CancelOrderController struct {
	cancelOrderPort in.CancelOrderPort
}

func NewCancelOrderController(cancelOrderPort in.CancelOrderPort) *CancelOrderController {
	return &CancelOrderController{cancelOrderPort: cancelOrderPort}
}

func (c *CancelOrderController) BindHttpCall(r *gin.Engine) {
	r.PUT("/order/:orderId/cancel", c.CancelOrder)
}

func (c *CancelOrderController) CancelOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	err := c.cancelOrderPort.CancelOrder(orderId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to cancel order"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Order canceled successfully"})
}
