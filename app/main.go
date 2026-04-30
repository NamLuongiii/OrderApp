package main

import (
	"OrderApp/auth"
	"OrderApp/common/postgresql"
	"OrderApp/inventory"
	"OrderApp/notification/gmail"
	"OrderApp/order"

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
