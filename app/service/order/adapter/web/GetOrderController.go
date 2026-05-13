package web

import (
	"OrderApp/service/order/application/port/in"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetOrderController struct {
	getOrderPort in.GetOrderPort
}

func NewGetOrderController(getOrderPort in.GetOrderPort) *GetOrderController {
	return &GetOrderController{
		getOrderPort: getOrderPort,
	}
}

func (c *GetOrderController) BindHttpCall(r *gin.Engine) {
	r.GET("/order/:orderId", c.GetOrder)
	r.GET("/orders", c.GetPaginatedOrders)
}

func (c *GetOrderController) GetOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	order, e := c.getOrderPort.GetOrder(orderId)
	if e != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve order"})
		return
	}

	orderResponse, e := mapOrder(order)
	if e != nil {
		ctx.JSON(500, gin.H{"error": "Failed to map order"})
		return
	}

	ctx.JSON(200, orderResponse)
}

func (c *GetOrderController) GetPaginatedOrders(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "0")
	limit := ctx.DefaultQuery("limit", "10")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid page number"})
		return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit number"})
		return
	}
	orders, e := c.getOrderPort.GetPaginatedOrders(intPage, intLimit)
	if e != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	orderResponses := make([]orderResponse, len(orders))
	for i, order := range orders {
		orderResponse, e := mapOrder(order)
		if e != nil {
			ctx.JSON(500, gin.H{"error": "Failed to map order"})
			return
		}
		orderResponses[i] = *orderResponse
	}

	ctx.JSON(200, orderResponses)
}
