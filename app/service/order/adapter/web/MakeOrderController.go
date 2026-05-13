package web

import (
	"OrderApp/service/order/application/domain/usecase"
	"OrderApp/service/order/application/port/in"

	"github.com/gin-gonic/gin"
)

type MakeOrderController struct {
	makeOrderPort in.MakeOrderPort
}

func NewMakeOrderController(makeOrderPort in.MakeOrderPort) *MakeOrderController {
	return &MakeOrderController{makeOrderPort: makeOrderPort}
}

func (c *MakeOrderController) BindHttpCall(r *gin.Engine) {
	r.POST("/order", c.MakeOrder)
}

type Body struct {
	Products []struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	} `json:"products"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"required,min=10,max=15"`
	Address string `json:"address" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Note    string `json:"note"`
}

func (c *MakeOrderController) MakeOrder(ctx *gin.Context) {
	body := Body{}

	e := ctx.ShouldBindBodyWithJSON(&body)
	if e != nil {
		ctx.JSON(400, gin.H{"msg": "Invalid request body"})
		return
	}

	makeOrderCommand, e := usecase.NewMakeOrderCommand(
		[]struct {
			ProductID string
			Quantity  int
		}(body.Products),
		body.Email,
		body.Phone,
		body.Address,
		body.Name,
		body.Note,
	)

	if e != nil {
		ctx.JSON(400, gin.H{"msg": e.Error()})
		return
	}

	e = c.makeOrderPort.MakeOrder(*makeOrderCommand)

	if e != nil {
		ctx.JSON(500, gin.H{"msg": e.Error()})
		return
	}

	ctx.JSON(200, body)
}
