package usecase

import "errors"

type MakeOrderCommand struct {
	Products []struct {
		ProductID string
		Quantity  int
	}
	Email   string
	Phone   string
	Address string
	Name    string
	Note    string
}

func NewMakeOrderCommand(
	products []struct {
		ProductID string
		Quantity  int
	},
	email,
	phone,
	address,
	name,
	note string) (*MakeOrderCommand, error) {
	if len(products) == 0 {
		return nil, errors.New("products must be at least 1")
	}
	return &MakeOrderCommand{
		Products: products,
		Email:    email,
		Phone:    phone,
		Address:  address,
		Name:     name,
		Note:     note,
	}, nil
}
