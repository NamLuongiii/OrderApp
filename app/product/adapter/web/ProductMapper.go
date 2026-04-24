package web

import "OrderApp/product/application/domain/model"

type ProductResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	SalePrice  string `json:"salePrice"`
	FinalPrice string `json:"finalPrice"`
}

func MapProduct(product *model.Product) *ProductResponse {
	if product == nil {
		return nil
	}

	salePrice := ""
	if product.GetSalePrice() != nil {
		salePrice = (*product.GetSalePrice()).String()
	}

	return &ProductResponse{
		ID:         product.GetId(),
		Name:       product.GetName(),
		Price:      product.GetPrice().String(),
		SalePrice:  salePrice,
		FinalPrice: product.GetFinalPrice().String(),
	}
}

func MapProducts(products []*model.Product) []*ProductResponse {
	productResponse := make([]*ProductResponse, len(products))
	for i, product := range products {
		productResponse[i] = MapProduct(product)
	}
	return productResponse
}
