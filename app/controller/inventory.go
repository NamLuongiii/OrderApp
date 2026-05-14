package controller

import (
	"OrderApp/common/msg"
	"OrderApp/service/inventory"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	service inventory.Service
	r       *gin.Engine
}

func NewInventoryController(r *gin.Engine, service inventory.Service) *InventoryController {
	return &InventoryController{r: r, service: service}
}

func (c *InventoryController) Init() {
	r := c.r
	g := r.Group("v1")
	g.GET("/product/:id", c.GetProduct)
	g.GET("/products", c.GetProductPagination)
	g.POST("/product", c.CreateProduct)

}

func (c *InventoryController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, e := c.service.GetProduct(id)

	if e != nil {
		ctx.JSON(http.StatusNotFound, e.Error())
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, msg.ProductNotFound)
		return
	}

	response := ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		SalePrice:  *product.SalePrice,
		FinalPrice: product.GetFinalPrice(),
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *InventoryController) GetProductPagination(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")

	var numPage = 1
	var numSize = 12

	if page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			numPage = p
		}
	}

	if size != "" {
		if s, err := strconv.Atoi(size); err == nil && s > 0 {
			numSize = s
		}
	}

	products, pagination, e := c.service.GetProductPagination(numPage, numSize)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	var response []ProductResponse
	for _, product := range products {
		response = append(response, ProductResponse{
			ID:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			SalePrice:  *product.SalePrice,
			FinalPrice: product.GetFinalPrice(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products":   response,
		"pagination": pagination,
	})
}

func (c *InventoryController) CreateProduct(ctx *gin.Context) {
	request := CreateProductRequest{}
	e := ctx.ShouldBindJSON(&request)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	id, e := c.service.CreateProduct(inventory.CreateProductCommand{
		Name:      request.Name,
		Price:     request.Price,
		SalePrice: request.SalePrice,
	})

	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}

type ProductResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	SalePrice  int64  `json:"salePrice"`
	FinalPrice int64  `json:"finalPrice"`
}

type CreateProductRequest struct {
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	SalePrice int64  `json:"salePrice"`
}
