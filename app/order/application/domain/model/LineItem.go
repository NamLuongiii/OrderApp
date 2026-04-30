package model

import "OrderApp/common/class"

type LineItem interface {
	GetID() *string
	GetProductName() string
	GetProductPrice() class.Money
	GetProductQuantity() int
	GetProductTotal() class.Money
	GetProductID() string
}

type lineItemImpl struct {
	id              *string
	productID       string
	productPrice    class.Money
	productQuantity int
	productTotal    class.Money
	productName     string
}

func NewLineItem(
	id *string,
	productID string,
	productPrice class.Money,
	productQuantity int,
	productTotal class.Money,
	productName string,
) LineItem {
	return &lineItemImpl{
		id:              id,
		productID:       productID,
		productPrice:    productPrice,
		productQuantity: productQuantity,
		productTotal:    productTotal,
		productName:     productName,
	}
}

func (l *lineItemImpl) GetID() *string               { return l.id }
func (l *lineItemImpl) GetProductName() string       { return l.productName }
func (l *lineItemImpl) GetProductPrice() class.Money { return l.productPrice }
func (l *lineItemImpl) GetProductQuantity() int      { return l.productQuantity }
func (l *lineItemImpl) GetProductTotal() class.Money { return l.productTotal }
func (l *lineItemImpl) GetProductID() string         { return l.productID }
