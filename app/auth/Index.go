package auth

import (
	"OrderApp/auth/adapter/persistence"
	"OrderApp/auth/adapter/web"
	"OrderApp/auth/application/domain/usecase"
	"OrderApp/auth/application/port/in"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boostrap(r *gin.Engine, db *gorm.DB) in.CheckRolePort {
	userPersistence := persistence.NewUserPersistence(db)

	createUser := usecase.NewCreateUserUseCase(userPersistence)

	createUserController := web.NewCreateUserController(createUser)
	createUserController.BindHttpCall(r)

	getUser := usecase.NewGetUser(userPersistence)
	getUserController := web.NewGetUserController(getUser)
	getUserController.BindHttpCall(r)

	login := usecase.NewLoginUseCase(userPersistence)
	web.NewLoginController(login).BindHttpCall(r)

	checkRole := usecase.NewCheckRoleMiddleware()
	return checkRole
}
