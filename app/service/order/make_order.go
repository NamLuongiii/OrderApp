package order

import "OrderApp/persistency/table"

func (i ServiceImpl) MakeOrder(command MakeOrderCommand) (string, error) {
	numberOfProducts := len(command.Products)
	extractIds := make([]string, numberOfProducts)
	for i, product := range command.Products {
		extractIds[i] = product.ID
	}

	products, e := i.productPersistency.GetProductsByIDs(extractIds)
	if e != nil {
		return "", e
	}

	var orderTotal int64 = 0
	for i, product := range products {
		quantity := command.Products[i].Quantity
		itemTotal := product.GetFinalPrice() * quantity
		orderTotal += itemTotal
	}

	var order = table.Order{
		Name:    command.Customer.Name,
		Total:   orderTotal,
		Email:   command.Customer.Email,
		Phone:   command.Customer.Phone,
		Address: command.Customer.Address,
		Note:    command.Customer.Note,
		Status:  string(StatusPending),
	}
	orderId, e := i.orderPersistency.SaveOrder(order)
	if e != nil {
		return "", e
	}

	lineItems := make([]*table.LineItem, numberOfProducts)
	for i, _ := range lineItems {
		lineItems[i] = &table.LineItem{
			OrderID:     orderId,
			ProductID:   products[i].ID,
			Quantity:    command.Products[i].Quantity,
			Price:       products[i].GetFinalPrice(),
			Total:       products[i].GetFinalPrice() * command.Products[i].Quantity,
			ProductName: products[i].Name,
		}
	}

	e = i.lineItemPersistency.SaveLineItems(lineItems)
	if e != nil {
		return "", e
	}

	return orderId, nil
}

type MakeOrderCommand struct {
	Products []struct {
		ID       string
		Quantity int64
	}
	Customer struct {
		Name    string
		Email   string
		Phone   string
		Address string
		Note    string
	}
}
