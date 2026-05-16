package controller

import (
	"OrderApp/common/class"
	"OrderApp/service/order"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	r            *gin.Engine
	orderService order.Service
}

func NewOrderController(r *gin.Engine, orderService order.Service) *OrderController {
	return &OrderController{r: r, orderService: orderService}
}

func (c *OrderController) Init() {
	r := c.r
	g := r.Group("v1")
	g.POST("/order", c.CreateOrder)
	g.GET("/order/:id", c.GetOrder)
	g.PUT("/order/:id", c.UpdateOrderStatus)
	g.GET("/orders", c.GetOrderPagination)
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	request := CreateOrderRequest{}
	e := ctx.ShouldBindJSON(&request)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	command := order.MakeOrderCommand{
		Products: []struct {
			ID       string
			Quantity int64
		}(request.Products),
		Customer: struct {
			Name    string
			Email   string
			Phone   string
			Address string
			Note    string
		}(request.Customer),
	}

	id, e := c.orderService.MakeOrder(command)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}
	ctx.JSON(http.StatusOK, id)
}

func (c *OrderController) GetOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	tableOrder, e := c.orderService.GetOrder(id)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	itemResponse := make([]ItemResponse, len(tableOrder.LineItems))
	for i, item := range tableOrder.LineItems {
		itemResponse[i] = ItemResponse{
			ID:        item.ID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Total:     item.Total,
			Name:      item.ProductName,
			ProductID: item.ProductID,
		}
	}
	response := OrderResponse{
		ID:        tableOrder.ID,
		Total:     tableOrder.Total,
		CreatedAt: tableOrder.CreatedAt.String(),
		UpdatedAt: tableOrder.UpdatedAt.String(),
		Status:    tableOrder.Status,
		Items:     itemResponse,
		Name:      tableOrder.Name,
		Email:     tableOrder.Email,
		Phone:     tableOrder.Phone,
		Address:   tableOrder.Address,
		Note:      tableOrder.Note,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	status := ctx.Query("status")
	id := ctx.Param("id")

	e := class.ValidateOrderStatus(status)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	e = c.orderService.UpdateOrderStatus(id, class.OrderStatus(status))
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}
	ctx.JSON(http.StatusOK, "success")
}

func (c *OrderController) GetOrderPagination(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")

	intPage := 1
	intPageSize := 12

	if page != "" {
		if p, e := strconv.Atoi(page); e == nil && p > 0 {
			intPage = p
		}
	}

	if size != "" {
		if s, e := strconv.Atoi(size); e == nil && s > 0 {
			intPageSize = s
		}
	}
	orders, pagination, e := c.orderService.GetOrderPagination(intPage, intPageSize)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	var response []OrderResponse
	for _, tableOrder := range orders {
		itemResponse := make([]ItemResponse, len(tableOrder.LineItems))
		for i, item := range tableOrder.LineItems {
			itemResponse[i] = ItemResponse{
				ID:        item.ID,
				Quantity:  item.Quantity,
				Price:     item.Price,
				Total:     item.Total,
				Name:      item.ProductName,
				ProductID: item.ID,
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"orders":     response,
		"pagination": pagination,
	})
}

type CreateOrderRequest struct {
	Products []struct {
		ID       string `json:"productId"`
		Quantity int64  `json:"quantity"`
	} `json:"items"`

	Customer struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		Note    string `json:"note"`
	} `json:"customer"`
}

type ItemResponse struct {
	ID        string `json:"itemID"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
	Total     int64  `json:"total"`
	Name      string `json:"name"`
	ProductID string `json:"productID"`
}

type OrderResponse struct {
	ID        string `json:"id"`
	Total     int64  `json:"total"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Note      string `json:"note"`

	Items []ItemResponse `json:"items"`
}
