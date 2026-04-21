package product

import (
	"OrderApp/product/adapter/persistence"
	"OrderApp/product/adapter/web"
	"OrderApp/product/application/domain/usecase"
	"OrderApp/product/application/port/in"
)

func Start() {
	persistenceAdapter := persistence.NewSaveProductPortImpl()
	createProductUseCase := usecase.NewCreateProduct(persistenceAdapter)
	createProductController := web.NewCreateProductController(createProductUseCase)

	command, e := in.NewCreateProductCommand("MacBook Pro", 100)
	if e != nil {
		panic(e)
		return
	}

	e = createProductController.CreateProduct(*command)
	if e != nil {
		panic(e)
		return
	}
}
