package in

import "errors"

type CreateProductCommand struct {
	Name  string
	Price int
}

func NewCreateProductCommand(name string, price int) (*CreateProductCommand, error) {
	e := validateName(name)
	if e != nil {
		return nil, e
	}

	return &CreateProductCommand{
		Name:  name,
		Price: price,
	}, nil
}

func validateName(name string) error {
	if len(name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if len(name) > 255 {
		return errors.New("name must be at most 255 characters")
	}

	return nil
}
