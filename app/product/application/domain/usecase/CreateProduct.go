package usecase

import (
	"OrderApp/common"
	"OrderApp/product/application/domain/model"
	"OrderApp/product/application/port/in"
	"OrderApp/product/application/port/out"
)

type CreateProduct struct {
	SaveProductPort out.PersistenceProductPort
}

func NewCreateProduct(saveProductPort out.PersistenceProductPort) *CreateProduct {
	return &CreateProduct{
		SaveProductPort: saveProductPort,
	}
}

func (service *CreateProduct) CreateProduct(command in.CreateProductCommand) error {
	price, e := common.NewPositiveMoney(command.Price)
	if e != nil {
		return e
	}

	product := model.NewProductWithoutId(command.Name, price, nil)
	return service.SaveProductPort.SaveProduct(product)
}
