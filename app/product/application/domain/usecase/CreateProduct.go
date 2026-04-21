package usecase

import (
	"OrderApp/product/application/port/in"
	"OrderApp/product/application/port/out"
	"fmt"
)

type CreateProduct struct {
	SaveProductPort out.SaveProductPort
}

func NewCreateProduct(saveProductPort out.SaveProductPort) *CreateProduct {
	return &CreateProduct{
		SaveProductPort: saveProductPort,
	}
}

func (service *CreateProduct) CreateProduct(command in.CreateProductCommand) error {
	fmt.Printf("CreateProduct %v\n", command.Name)
	return service.SaveProductPort.SaveProduct()
}
