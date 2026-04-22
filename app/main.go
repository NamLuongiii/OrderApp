package main

import (
	"OrderApp/product"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	product.Boostrap(r)
	e := r.Run(":8080")
	e.Error()
}
