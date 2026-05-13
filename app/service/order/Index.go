package order

import (
	"OrderApp/service/inventory/application/port/in"
	"OrderApp/service/notification/port"
	persistence2 "OrderApp/service/order/adapter/persistence"
	web2 "OrderApp/service/order/adapter/web"
	usecase2 "OrderApp/service/order/application/domain/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boostrap(
	r *gin.Engine,
	db *gorm.DB,
	inventoryPort in.InventoryPort,
	mailServicePort port.MailServicePort,
) {
	orderPersistence := persistence2.NewOrderPersistence(db)
	itemLinePersistence := persistence2.NewLineItemPersistence(db)

	makeOrderUseCase := usecase2.NewMakeOrder(
		inventoryPort,
		orderPersistence,
		itemLinePersistence,
		mailServicePort,
	)

	makeOrderController := web2.NewMakeOrderController(makeOrderUseCase)

	getOrderUseCase := usecase2.NewGetOrder(orderPersistence)
	getOrderController := web2.NewGetOrderController(getOrderUseCase)

	cancelOrderUseCase := usecase2.NewCancelOrder(orderPersistence)
	web2.NewCancelOrderController(cancelOrderUseCase).BindHttpCall(r)

	confirmOrderUseCase := usecase2.NewConfirmOrder(orderPersistence)
	web2.NewConfirmOrderController(confirmOrderUseCase).BindHttpCall(r)

	markOrderCompletedUseCase := usecase2.NewMarkOrderCompletedUseCase(orderPersistence)
	web2.NewMarkOrderCompletedController(markOrderCompletedUseCase).BindHttpCall(r)

	markOrderDeliveredUseCase := usecase2.NewMarkOrderDeliveredUseCase(orderPersistence)
	web2.NewMarkOrderDeliveredController(markOrderDeliveredUseCase).BindHttpCall(r)

	makeOrderController.BindHttpCall(r)
	getOrderController.BindHttpCall(r)
}
