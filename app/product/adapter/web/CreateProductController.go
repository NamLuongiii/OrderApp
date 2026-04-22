package web

import (
	"OrderApp/product/application/port/in"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	CreateProductUseCase in.CreateProductUseCase
}

func NewCreateProductController(createProductUseCase in.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		CreateProductUseCase: createProductUseCase,
	}
}

func (c *CreateProductController) CreateProduct(httpCtx *gin.Context) {
	command := in.CreateProductCommand{
		Name:  "MacBook pro",
		Price: 100,
	}

	e := c.CreateProductUseCase.CreateProduct(command)

	if e != nil {
		httpCtx.JSON(500, gin.H{
			"message": e.Error(),
		})
	} else {
		httpCtx.JSON(200, gin.H{
			"message": "success",
		})
	}
}
