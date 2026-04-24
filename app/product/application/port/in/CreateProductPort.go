package in

type CreateProductPort interface {
	CreateProduct(command CreateProductCommand) error
}
