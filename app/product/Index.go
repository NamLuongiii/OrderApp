package product

import (
	"OrderApp/product/adapter/persistence"
	"OrderApp/product/adapter/web"
	"OrderApp/product/application/domain/usecase"

	"github.com/gin-gonic/gin"
)

func Boostrap(r *gin.Engine) {
	persistenceAdapter := persistence.NewSaveProductPortImpl()
	createProductUseCase := usecase.NewCreateProduct(persistenceAdapter)
	createProductController := web.NewCreateProductController(createProductUseCase)

	r.POST("/product", createProductController.CreateProduct)
}
