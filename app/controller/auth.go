package controller

import (
	"OrderApp/common/msg"
	"OrderApp/service/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	r           *gin.Engine
	authService auth.Service
}

func NewAuthController(r *gin.Engine, service auth.Service) *AuthController {
	return &AuthController{r: r, authService: service}
}

func (a *AuthController) Init() {
	r := a.r
	g := r.Group("v1")
	g.GET("/me", auth.CheckRole(), a.getUser)
	g.POST("/user", a.createUser)
	g.POST("/login", a.login)
}

func (a *AuthController) getUser(ctx *gin.Context) {
	id := ctx.GetString("userId")
	user, e := a.authService.GetUser(id)

	if e != nil {
		ctx.JSON(404, msg.UserNotFound)
		return
	}

	userResponse := GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      user.Role,
	}

	ctx.JSON(200, userResponse)
}

func (a *AuthController) createUser(ctx *gin.Context) {
	createUserRequest := CreateUserRequest{}
	e := ctx.ShouldBindJSON(&createUserRequest)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}
	command := auth.CreateUserCommand{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Role:     string(createUserRequest.Role),
		Password: createUserRequest.Password,
	}
	e = a.authService.CreateUser(command)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}
	ctx.JSON(http.StatusOK, 1)
}

func (a *AuthController) login(ctx *gin.Context) {
	loginRequest := LoginRequest{}
	e := ctx.ShouldBindJSON(&loginRequest)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	token, e := a.authService.Login(loginRequest.Email, loginRequest.Password)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, token)
}

type GetUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     auth.Role `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
