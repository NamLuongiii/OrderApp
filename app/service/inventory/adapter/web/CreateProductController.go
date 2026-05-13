package web

import (
	authService "OrderApp/service/auth/application/port/in"
	in2 "OrderApp/service/inventory/application/port/in"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	CreateProductUseCase in2.CreateProductPort
	CheckRolePort        authService.CheckRolePort
}

func NewCreateProductController(
	createProductUseCase in2.CreateProductPort,
	checkRolePort authService.CheckRolePort,
) *CreateProductController {
	return &CreateProductController{
		CreateProductUseCase: createProductUseCase,
		CheckRolePort:        checkRolePort,
	}
}

func (c *CreateProductController) BindHttpCall(r *gin.Engine) {
	r.POST("/product", c.CheckRolePort.CheckRole(), c.CreateProduct)
}

func (c *CreateProductController) CreateProduct(httpCtx *gin.Context) {
	body := struct {
		Name  string
		Price string
	}{}

	e := httpCtx.ShouldBindJSON(&body)
	if e != nil {
		returnError(httpCtx, 400, e)
		return
	}

	command, e := in2.NewCreateProductCommand(body.Name, body.Price)
	if e != nil {
		returnError(httpCtx, 400, e)
		return
	}

	e = c.CreateProductUseCase.CreateProduct(*command)
	if e != nil {
		returnError(httpCtx, 500, e)
		return
	}

	httpCtx.JSON(200, gin.H{
		"message": "success",
	})
}

func returnError(httpCtx *gin.Context, errorCode int, e error) {
	httpCtx.JSON(errorCode, gin.H{
		"message": e.Error(),
	})
}
