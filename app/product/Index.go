package product

import (
	"OrderApp/product/adapter/persistence"
	"OrderApp/product/adapter/web"
	"OrderApp/product/application/domain/usecase"

	"github.com/gin-gonic/gin"
)

func Boostrap(r *gin.Engine) {
	persistenceProductAdapter := persistence.ProductAdapterImpl()

	createProductUseCase := usecase.NewCreateProduct(persistenceProductAdapter)
	createProductController := web.NewCreateProductController(createProductUseCase)
	createProductController.BindHttpCall(r)

	getProductUseCase := usecase.NewGetProduct(persistenceProductAdapter)
	getProductController := web.NewGetProductController(getProductUseCase)
	getProductController.BindHttpCall(r)
}
