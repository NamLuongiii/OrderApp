package usecase

import (
	productModel "OrderApp/service/inventory/application/domain/model"
	"OrderApp/service/inventory/application/port/in"
	mailServiceModel "OrderApp/service/notification/gmail/model"
	"OrderApp/service/notification/port"
	model2 "OrderApp/service/order/application/domain/model"
	out2 "OrderApp/service/order/application/port/out"
	"errors"
)

type MakeOrder struct {
	inventoryPort           in.InventoryPort
	orderPersistencePort    out2.OrderPersistencePort
	lineItemPersistencePort out2.LineItemPersistencePort
	mailServicePort         port.MailServicePort
}

func NewMakeOrder(
	inventoryPort in.InventoryPort,
	orderPersistencePort out2.OrderPersistencePort,
	lineItemPersistencePort out2.LineItemPersistencePort,
	mailServicePort port.MailServicePort,
) *MakeOrder {
	return &MakeOrder{
		inventoryPort:           inventoryPort,
		orderPersistencePort:    orderPersistencePort,
		lineItemPersistencePort: lineItemPersistencePort,
		mailServicePort:         mailServicePort,
	}
}

func (service *MakeOrder) MakeOrder(command MakeOrderCommand) error {
	lineItems, e := service.CreateLineItems(command)
	if e != nil {
		return e
	}

	order := model2.NewOrderWithoutId(
		lineItems,
		command.Email,
		command.Phone,
		command.Address,
		command.Name,
		command.Note,
		model2.PROCESSING,
	)

	orderId, e := service.orderPersistencePort.SaveOrder(order)
	if e != nil {
		return e
	}

	err := service.lineItemPersistencePort.SaveLineItems(orderId, lineItems)
	if err != nil {
		return err
	}

	go service.SendOrderConfirmationEmail(orderId, order)

	return nil
}

func (service *MakeOrder) CreateLineItems(command MakeOrderCommand) ([]model2.LineItem, error) {
	productIdArray := make([]string, len(command.Products))
	for i, product := range command.Products {
		productIdArray[i] = product.ProductID
	}

	productArray, e := service.inventoryPort.GetProductsBatch(productIdArray)
	if e != nil {
		return nil, e
	}

	productMap := make(map[string]*productModel.Product, len(productArray))
	for _, product := range productArray {
		productMap[product.GetId()] = product
	}

	lineItems := make([]model2.LineItem, len(command.Products))
	for i, item := range command.Products {
		product := productMap[item.ProductID]
		if product == nil {
			return nil, errors.New("product not found")
		}
		lineItems[i] = model2.NewLineItem(
			nil,
			product.GetId(),
			product.GetPrice(),
			item.Quantity,
			product.GetFinalPrice().Multiple(float64(item.Quantity)),
			product.GetName(),
		)
	}

	return lineItems, nil
}

func (service *MakeOrder) SendOrderConfirmationEmail(orderId string, order model2.Order) {
	products := make([]mailServiceModel.ProductData, len(order.GetLineItems()))
	for i, item := range order.GetLineItems() {
		products[i] = mailServiceModel.ProductData{
			ID:          item.GetProductID(),
			Quantity:    item.GetProductQuantity(),
			ProductName: item.GetProductName(),
			Price:       item.GetProductPrice().String(),
			Total:       item.GetProductTotal().String(),
		}
	}

	service.mailServicePort.SendSuccessfullyOrderPlayed(
		order.GetEmail(),
		mailServiceModel.SuccessfullyOrderPlayedData{
			OrderID:  orderId,
			Products: products,
		})
}
