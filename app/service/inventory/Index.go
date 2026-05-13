package inventory

import (
	authService "OrderApp/service/auth/application/port/in"
	"OrderApp/service/inventory/adapter/persistence"
	web2 "OrderApp/service/inventory/adapter/web"
	usecase2 "OrderApp/service/inventory/application/domain/usecase"
	"OrderApp/service/inventory/application/port/in"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boostrap(r *gin.Engine, db *gorm.DB, checkRolePort authService.CheckRolePort) in.InventoryPort {
	persistenceProductAdapter := persistence.ProductAdapterImpl(db)

	createProductUseCase := usecase2.NewCreateProduct(persistenceProductAdapter)
	createProductController := web2.NewCreateProductController(createProductUseCase, checkRolePort)
	createProductController.BindHttpCall(r)

	getProductUseCase := usecase2.NewGetProduct(persistenceProductAdapter)
	getProductController := web2.NewGetProductController(getProductUseCase)
	getProductController.BindHttpCall(r)

	return usecase2.NewInventoryPort(persistenceProductAdapter)
}
