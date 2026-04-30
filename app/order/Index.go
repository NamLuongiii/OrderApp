package order

import (
	"OrderApp/inventory/application/port/in"
	"OrderApp/notification/port"
	"OrderApp/order/adapter/persistence"
	"OrderApp/order/adapter/web"
	"OrderApp/order/application/domain/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boostrap(
	r *gin.Engine,
	db *gorm.DB,
	inventoryPort in.InventoryPort,
	mailServicePort port.MailServicePort,
) {
	orderPersistence := persistence.NewOrderPersistence(db)
	itemLinePersistence := persistence.NewLineItemPersistence(db)

	makeOrderUseCase := usecase.NewMakeOrder(
		inventoryPort,
		orderPersistence,
		itemLinePersistence,
		mailServicePort,
	)

	makeOrderController := web.NewMakeOrderController(makeOrderUseCase)

	getOrderUseCase := usecase.NewGetOrder(orderPersistence)
	getOrderController := web.NewGetOrderController(getOrderUseCase)

	cancelOrderUseCase := usecase.NewCancelOrder(orderPersistence)
	web.NewCancelOrderController(cancelOrderUseCase).BindHttpCall(r)

	confirmOrderUseCase := usecase.NewConfirmOrder(orderPersistence)
	web.NewConfirmOrderController(confirmOrderUseCase).BindHttpCall(r)

	markOrderCompletedUseCase := usecase.NewMarkOrderCompletedUseCase(orderPersistence)
	web.NewMarkOrderCompletedController(markOrderCompletedUseCase).BindHttpCall(r)

	markOrderDeliveredUseCase := usecase.NewMarkOrderDeliveredUseCase(orderPersistence)
	web.NewMarkOrderDeliveredController(markOrderDeliveredUseCase).BindHttpCall(r)

	makeOrderController.BindHttpCall(r)
	getOrderController.BindHttpCall(r)
}
