package main

import (
	"OrderApp/common/cors"
	"OrderApp/common/postgresql"
	"OrderApp/service/auth"
	"OrderApp/service/inventory"
	"OrderApp/service/notification/gmail"
	"OrderApp/service/order"

	"github.com/gin-gonic/gin"
)

func main() {
	db, e := postgresql.NewConnection()
	if e != nil {
		panic(e)
		return
	}

	mailService, e := gmail.NewMailService()
	if e != nil {
		panic(e)
		return
	}

	r := gin.Default()

	r.Use(cors.MiddlewareCors())

	checkRolePort := auth.Boostrap(r, db)
	inventoryPort := inventory.Boostrap(r, db, checkRolePort)
	order.Boostrap(
		r,
		db,
		inventoryPort,
		mailService,
	)

	e = r.Run(":8080")
	if e != nil {
		panic(e)
	}
}
