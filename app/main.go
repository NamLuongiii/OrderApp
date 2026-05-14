package main

import (
	"OrderApp/common/cors"
	"OrderApp/common/postgresql"
	"OrderApp/controller"
	"OrderApp/persistency"
	"OrderApp/service/auth"
	"OrderApp/service/inventory"
	"OrderApp/service/mail"
	"OrderApp/service/order"

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

	mailService, e := mail.NewMailService()
	if e != nil {
		panic(e)
		return
	}

	userPersistency := persistency.NewUserPersistence(db)
	authService := auth.NewService(userPersistency)
	authController := controller.NewAuthController(r, authService)
	authController.Init()

	productPersistency := persistency.NewProductPersistency(db)
	inventoryService := inventory.NewInventoryService(productPersistency)
	inventoryController := controller.NewInventoryController(r, inventoryService)
	inventoryController.Init()

	orderPersistency := persistency.NewOrderPersistence(db)
	orderService := order.NewService(orderPersistency,
		persistency.NewLineItemPersistence(db),
		productPersistency,
		mailService)
	orderController := controller.NewOrderController(r, orderService)
	orderController.Init()

	e = r.Run(":8080")
	if e != nil {
		panic(e)
	}
}
