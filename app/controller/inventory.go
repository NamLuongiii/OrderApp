package controller

import (
	"OrderApp/common/class"
	"OrderApp/common/msg"
	"OrderApp/service/inventory"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	service inventory.Service
}

func NewInventoryController(service inventory.Service) *InventoryController {
	return &InventoryController{service: service}
}

func (c *InventoryController) Init(r *gin.Engine) {
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

	price, e := class.NewPositiveMoney(request.Price)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
	}

	salePrice, e := class.NewPositiveMoney(request.SalePrice)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
	}

	id, e := c.service.CreateProduct(inventory.CreateProductCommand{
		Name:      request.Name,
		Price:     price,
		SalePrice: &salePrice,
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
	Price      string `json:"price"`
	SalePrice  string `json:"salePrice"`
	FinalPrice string `json:"finalPrice"`
}

type CreateProductRequest struct {
	Name      string `json:"name"`
	Price     string `json:"price"`
	SalePrice string `json:"salePrice"`
}
