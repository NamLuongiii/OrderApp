package persistence

import (
	"OrderApp/common/class"
	"OrderApp/common/postgresql/table"
	model2 "OrderApp/service/order/application/domain/model"
)

func MapOrder(persistenceOrder table.Order) (model2.Order, error) {
	lineItems := make([]model2.LineItem, len(persistenceOrder.LineItems))
	for i, item := range persistenceOrder.LineItems {
		price, e := class.NewPositiveMoney(item.Price)
		if e != nil {
			return nil, e
		}
		totalPrice, e := class.NewPositiveMoney(item.Total)
		if e != nil {
			return nil, e
		}
		lineItems[i] = model2.NewLineItem(
			&item.ID,
			item.ProductID,
			price,
			item.Quantity,
			totalPrice,
			item.ProductName,
		)
	}

	total, e := class.NewPositiveMoney(persistenceOrder.Total)
	if e != nil {
		return nil, e
	}

	order := model2.NewOrder(
		persistenceOrder.ID,
		lineItems,
		total,
		persistenceOrder.Email,
		persistenceOrder.Phone,
		persistenceOrder.Address,
		persistenceOrder.Name,
		persistenceOrder.Note,
		model2.Status(persistenceOrder.Status),
	)

	return order, nil
}
