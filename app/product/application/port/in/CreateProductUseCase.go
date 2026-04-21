package in

type CreateProductUseCase interface {
	CreateProduct(command CreateProductCommand) error
}
