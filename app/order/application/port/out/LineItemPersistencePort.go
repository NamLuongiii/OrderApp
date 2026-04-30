package out

import "OrderApp/order/application/domain/model"

type LineItemPersistencePort interface {
	SaveLineItems(orderId string, lineItems []model.LineItem) error
}
