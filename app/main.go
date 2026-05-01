package main

import (
	"OrderApp/auth"
	"OrderApp/common/postgresql"
	"OrderApp/inventory"
	"OrderApp/notification/gmail"
	"OrderApp/order"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware cấu hình các header để cho phép truy cập từ các domain khác
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Cho phép tất cả các nguồn
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

	// SỬ DỤNG MIDDLEWARE CORS TẠI ĐÂY
	r.Use(CORSMiddleware())

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
