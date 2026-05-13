package web

import (
	"OrderApp/service/auth/application/port/in"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	logPort in.LoginPort
}

func NewLoginController(logPort in.LoginPort) *LoginController {
	return &LoginController{logPort: logPort}
}

func (c *LoginController) BindHttpCall(r *gin.Engine) {
	r.POST("/login", c.Login)
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l LoginController) Login(ctx *gin.Context) {
	body := LoginBody{}
	e := ctx.ShouldBindJSON(&body)
	if e != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	token, e := l.logPort.Login(body.Email, body.Password)
	if e != nil {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}
