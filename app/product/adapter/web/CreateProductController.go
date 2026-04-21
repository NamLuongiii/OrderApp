package web

import (
	"OrderApp/product/application/port/in"
)

type CreateProductController struct {
	CreateProductUseCase in.CreateProductUseCase
}

func NewCreateProductController(createProductUseCase in.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		CreateProductUseCase: createProductUseCase,
	}
}

func (c *CreateProductController) CreateProduct(command in.CreateProductCommand) error {
	return c.CreateProductUseCase.CreateProduct(command)
}
