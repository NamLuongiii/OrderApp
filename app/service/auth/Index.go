package auth

import (
	"OrderApp/service/auth/adapter/persistence"
	web2 "OrderApp/service/auth/adapter/web"
	usecase2 "OrderApp/service/auth/application/domain/usecase"
	"OrderApp/service/auth/application/port/in"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boostrap(r *gin.Engine, db *gorm.DB) in.CheckRolePort {
	userPersistence := persistence.NewUserPersistence(db)

	createUser := usecase2.NewCreateUserUseCase(userPersistence)

	createUserController := web2.NewCreateUserController(createUser)
	createUserController.BindHttpCall(r)

	getUser := usecase2.NewGetUser(userPersistence)
	getUserController := web2.NewGetUserController(getUser)
	getUserController.BindHttpCall(r)

	login := usecase2.NewLoginUseCase(userPersistence)
	web2.NewLoginController(login).BindHttpCall(r)

	checkRole := usecase2.NewCheckRoleMiddleware()
	return checkRole
}
