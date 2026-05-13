package model

import (
	"OrderApp/common/class"
	"time"
)

type Order interface {
	GetID() string
	GetLineItems() []LineItem
	GetTotal() class.Money
	GetEmail() string
	GetPhone() string
	GetAddress() string
	GetName() string
	GetNote() string
	GetStatus() Status
	GetCreatedAt() int64
	GetUpdatedAt() int64
}

type orderImpl struct {
	id        *string
	lineItems []LineItem
	total     class.Money
	email     string
	phone     string
	address   string
	name      string
	note      string
	status    Status
	createdAt int64
	UpdatedAt int64
}

func NewOrder(
	id string,
	lineItems []LineItem,
	total class.Money,
	email string,
	phone string,
	address string,
	name string,
	note string,
	status Status,
) Order {
	return &orderImpl{
		id:        &id,
		lineItems: lineItems,
		total:     total,
		email:     email,
		phone:     phone,
		address:   address,
		name:      name,
		note:      note,
		status:    status,
		createdAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}

func NewOrderWithoutId(
	lineItems []LineItem,
	email string,
	phone string,
	address string,
	name string,
	note string,
	status Status,
) Order {
	total := calculateOrderTotal(lineItems)
	return &orderImpl{
		lineItems: lineItems,
		total:     total,
		email:     email,
		phone:     phone,
		address:   address,
		name:      name,
		note:      note,
		status:    status,
		createdAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}

func (o *orderImpl) GetID() string            { return *o.id }
func (o *orderImpl) GetLineItems() []LineItem { return o.lineItems }
func (o *orderImpl) GetTotal() class.Money    { return o.total }
func (o *orderImpl) GetEmail() string         { return o.email }
func (o *orderImpl) GetPhone() string         { return o.phone }
func (o *orderImpl) GetAddress() string       { return o.address }
func (o *orderImpl) GetName() string          { return o.name }
func (o *orderImpl) GetNote() string          { return o.note }
func (o *orderImpl) GetStatus() Status        { return o.status }
func (o *orderImpl) GetCreatedAt() int64      { return o.createdAt }
func (o *orderImpl) GetUpdatedAt() int64      { return o.UpdatedAt }

func calculateOrderTotal(lineItems []LineItem) class.Money {
	total, _ := class.NewMoney("0")
	for _, item := range lineItems {
		total = total.Add(item.GetProductTotal())
	}
	return total
}
