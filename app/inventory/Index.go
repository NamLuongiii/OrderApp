package inventory

import (
	authService "OrderApp/auth/application/port/in"
	"OrderApp/inventory/adapter/persistence"
	"OrderApp/inventory/adapter/web"
	"OrderApp/inventory/application/domain/usecase"
	"OrderApp/inventory/application/port/in"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boostrap(r *gin.Engine, db *gorm.DB, checkRolePort authService.CheckRolePort) in.InventoryPort {
	persistenceProductAdapter := persistence.ProductAdapterImpl(db)

	createProductUseCase := usecase.NewCreateProduct(persistenceProductAdapter)
	createProductController := web.NewCreateProductController(createProductUseCase, checkRolePort)
	createProductController.BindHttpCall(r)

	getProductUseCase := usecase.NewGetProduct(persistenceProductAdapter)
	getProductController := web.NewGetProductController(getProductUseCase)
	getProductController.BindHttpCall(r)

	return usecase.NewInventoryPort(persistenceProductAdapter)
}
