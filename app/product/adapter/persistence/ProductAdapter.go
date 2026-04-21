package persistence

import "fmt"

type ProductAdapter struct {
}

func NewSaveProductPortImpl() *ProductAdapter {
	return &ProductAdapter{}
}

func (p *ProductAdapter) SaveProduct() error {
	fmt.Println("Save product to database")
	return nil
}
