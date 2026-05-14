package order

import (
	"OrderApp/common/class"
	"OrderApp/common/msg"
	"OrderApp/persistency/table"
	"OrderApp/service/mail"
	"errors"
	"math"

	"github.com/shopspring/decimal"
)

func (s ServiceImpl) MakeOrder(command MakeOrderCommand) (string, error) {
	numberOfProducts := len(command.Products)
	extractIds := make([]string, numberOfProducts)
	for i, product := range command.Products {
		extractIds[i] = product.ID
	}

	products, e := s.productPersistency.GetProductsByIDs(extractIds)
	if e != nil {
		return "", e
	}

	var orderTotal decimal.Decimal = decimal.Zero
	maxSafeValue := decimal.NewFromInt(math.MaxInt64)

	for i, product := range products {
		quantity := command.Products[i].Quantity
		orderTotal := decimal.NewFromInt(product.GetFinalPrice()).Mul(decimal.NewFromInt(quantity))
		if orderTotal.GreaterThanOrEqual(maxSafeValue) {
			return "", errors.New(msg.PriceValueTooLarge)
		}
		orderTotal = orderTotal.Add(orderTotal)
	}

	if orderTotal.GreaterThan(maxSafeValue) {
		return "", errors.New(msg.PriceValueTooLarge)
	}

	var order = table.Order{
		Name:    command.Customer.Name,
		Total:   orderTotal.IntPart(),
		Email:   command.Customer.Email,
		Phone:   command.Customer.Phone,
		Address: command.Customer.Address,
		Note:    command.Customer.Note,
		Status:  string(class.StatusPending),
	}
	orderId, e := s.orderPersistency.SaveOrder(order)
	if e != nil {
		return "", e
	}

	lineItems := make([]*table.LineItem, numberOfProducts)
	for i, _ := range lineItems {
		itemTotal := decimal.NewFromInt(products[i].GetFinalPrice()).Mul(decimal.NewFromInt(command.Products[i].Quantity))
		if itemTotal.GreaterThan(maxSafeValue) {
			return "", errors.New(msg.PriceValueTooLarge)
		}

		lineItems[i] = &table.LineItem{
			OrderID:     orderId,
			ProductID:   products[i].ID,
			Quantity:    command.Products[i].Quantity,
			Price:       products[i].GetFinalPrice(),
			Total:       itemTotal.IntPart(),
			ProductName: products[i].Name,
		}
	}

	e = s.lineItemPersistency.SaveLineItems(lineItems)
	if e != nil {
		return "", e
	}

	mailProducts := make([]mail.ProductData, numberOfProducts)
	for i, _ := range mailProducts {
		mailProducts[i] = mail.ProductData{
			ID:          products[i].ID,
			Quantity:    command.Products[i].Quantity,
			Price:       products[i].GetFinalPrice(),
			ProductName: products[i].Name,
			Total:       lineItems[i].Total,
		}
	}
	go s.mailService.SendNewOrderPlayed(order.Email, mail.SendNewOrderPlayedCommand{
		OrderID:  orderId,
		Products: mailProducts,
	})

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
