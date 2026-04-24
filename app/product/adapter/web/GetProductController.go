package web

import (
	"OrderApp/product/application/port/in"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductController struct {
	GetProductPort in.GetProductPort
}

func NewGetProductController(getProductPort in.GetProductPort) *GetProductController {
	return &GetProductController{
		GetProductPort: getProductPort,
	}
}

func (c *GetProductController) BindHttpCall(r *gin.Engine) {
	r.GET("/product", c.handleGetProducts)
	r.GET("/product/:id", c.handleGetProduct)
}

func (c *GetProductController) handleGetProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	product, e := c.GetProductPort.GetProduct(productId)
	if e != nil {
		returnError(ctx, 500, e)
		return
	}

	ctx.JSON(200, MapProduct(product))
}

func (c *GetProductController) handleGetProducts(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "0")
	limit := ctx.DefaultQuery("limit", "10")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		returnError(ctx, 400, err)
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		returnError(ctx, 400, err)
		return
	}

	products, e := c.GetProductPort.GetPaginatedProducts(intPage, intLimit)
	if e != nil {
		returnError(ctx, 500, e)
		return
	}

	ctx.JSON(200, MapProducts(products))
}
