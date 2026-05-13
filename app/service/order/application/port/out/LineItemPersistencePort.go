package out

import (
	"OrderApp/service/order/application/domain/model"
)

type LineItemPersistencePort interface {
	SaveLineItems(orderId string, lineItems []model.LineItem) error
}
