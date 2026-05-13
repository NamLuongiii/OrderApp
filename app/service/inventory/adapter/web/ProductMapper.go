package web

import (
	"OrderApp/service/inventory/application/domain/model"
)

type productResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	SalePrice  string `json:"salePrice"`
	FinalPrice string `json:"finalPrice"`
}

func mapProduct(product *model.Product) *productResponse {
	if product == nil {
		return nil
	}

	salePrice := ""
	if product.GetSalePrice() != nil {
		salePrice = (*product.GetSalePrice()).String()
	}

	return &productResponse{
		ID:         product.GetId(),
		Name:       product.GetName(),
		Price:      product.GetPrice().String(),
		SalePrice:  salePrice,
		FinalPrice: product.GetFinalPrice().String(),
	}
}

func MapProducts(products []*model.Product) []*productResponse {
	productResponse := make([]*productResponse, len(products))
	for i, product := range products {
		productResponse[i] = mapProduct(product)
	}
	return productResponse
}
