package web

import "OrderApp/order/application/domain/model"

type lineItemResponse struct {
	ProductId   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
	Total       string `json:"total"`
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
}
type orderResponse struct {
	ID        string             `json:"id"`
	LineItems []lineItemResponse `json:"line_items"`
	Name      string             `json:"name"`
	Total     string             `json:"total"`
	Phone     string             `json:"phone"`
	Email     string             `json:"email"`
	Address   string             `json:"address"`
	Note      string             `json:"note"`
	Status    string             `json:"status"`
	CreatedAt int64              `json:"created_at"`
	UpdatedAt int64              `json:"updated_at"`
}

func mapOrder(order model.Order) (*orderResponse, error) {
	if order == nil {
		return nil, nil
	}

	var lineItems []lineItemResponse
	for _, item := range order.GetLineItems() {
		lineItems = append(lineItems, lineItemResponse{
			ProductId:   item.GetProductID(),
			Quantity:    item.GetProductQuantity(),
			Price:       item.GetProductPrice().String(),
			Total:       item.GetProductTotal().String(),
			ID:          *item.GetID(),
			ProductName: item.GetProductName(),
		})
	}

	return &orderResponse{
		ID:        order.GetID(),
		LineItems: lineItems,
		Total:     order.GetTotal().String(),
		Name:      order.GetName(),
		Phone:     order.GetPhone(),
		Email:     order.GetEmail(),
		Address:   order.GetAddress(),
		Note:      order.GetNote(),
		Status:    string(order.GetStatus()),
		CreatedAt: order.GetCreatedAt(),
		UpdatedAt: order.GetUpdatedAt(),
	}, nil
}
