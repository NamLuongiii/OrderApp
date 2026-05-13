package usecase

import (
	"OrderApp/common/class"
	"OrderApp/service/inventory/application/domain/model"
	"OrderApp/service/inventory/application/port/in"
	"OrderApp/service/inventory/application/port/out"
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
	price, e := class.NewPositiveMoney(command.Price)
	if e != nil {
		return e
	}

	product := model.NewProductWithoutId(command.Name, price, nil)
	return service.SaveProductPort.SaveProduct(product)
}
