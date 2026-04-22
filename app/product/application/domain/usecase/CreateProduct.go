package usecase

import (
	"OrderApp/common"
	"OrderApp/product/application/domain/model"
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
	m, _ := common.NewMoney("10000")
	sm, _ := common.NewMoney("9000")
	prod := model.NewProduct("1", "a", m, &sm)
	fmt.Printf("CreateProduct %v\n", prod.GetFinalPrice().String())
	fmt.Printf("CreateProduct %v\n", command.Name)
	return service.SaveProductPort.SaveProduct()
}
