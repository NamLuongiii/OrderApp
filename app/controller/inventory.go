package controller

import "github.com/gin-gonic/gin"

type InventoryController struct {
}

func NewInventoryController() *InventoryController {
	return &InventoryController{}
}

func (c *InventoryController) Init(r *gin.Engine) {

}
