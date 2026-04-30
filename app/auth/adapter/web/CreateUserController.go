package web

import (
	"OrderApp/auth/application/domain/model"
	"OrderApp/auth/application/port/in"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	createUserPort in.CreateUserPort
}

func NewCreateUserController(createUserPort in.CreateUserPort) *CreateUserController {
	return &CreateUserController{
		createUserPort: createUserPort,
	}
}

func (c *CreateUserController) BindHttpCall(r *gin.Engine) {
	r.POST("/user", c.CreateUser)
}

type Body struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *CreateUserController) CreateUser(ctx *gin.Context) {
	body := Body{}

	e := ctx.ShouldBindJSON(&body)
	if e != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	user, e := model.NewUserWithoutId(
		model.RoleAdmin,
		body.Email,
		body.Password,
		body.Name,
	)
	e = c.createUserPort.CreateUser(user)
	if e != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User created successfully"})
}
