package main

import (
	"OrderApp/common/cors"
	"OrderApp/common/postgresql"
	"OrderApp/controller"
	"OrderApp/persistency"
	"OrderApp/service/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	db, e := postgresql.NewConnection()
	if e != nil {
		panic(e)
		return
	}

	r := gin.Default()

	r.Use(cors.MiddlewareCors())

	userPersistency := persistency.NewUserPersistence(db)
	authService := auth.NewService(userPersistency)
	authController := controller.NewAuthController(r, authService)
	authController.Init()

	e = r.Run(":8080")
	if e != nil {
		panic(e)
	}
}
