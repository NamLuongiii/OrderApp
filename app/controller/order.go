package controller

import "github.com/gin-gonic/gin"

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (c *OrderController) Init(r *gin.Engine) {}
