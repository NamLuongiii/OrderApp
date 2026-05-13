package web

import (
	"OrderApp/service/order/application/port/in"

	"github.com/gin-gonic/gin"
)

type ConfirmOrderController struct {
	confirmOrderPort in.ConfirmOrderPort
}

func NewConfirmOrderController(confirmOrderPort in.ConfirmOrderPort) *ConfirmOrderController {
	return &ConfirmOrderController{confirmOrderPort: confirmOrderPort}
}

func (c *ConfirmOrderController) BindHttpCall(r *gin.Engine) {
	r.PUT("/order/:orderId/confirm", c.ConfirmOrder)
}

func (c *ConfirmOrderController) ConfirmOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	err := c.confirmOrderPort.ConfirmOrder(orderId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to confirm order"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Order confirmed successfully"})
}
