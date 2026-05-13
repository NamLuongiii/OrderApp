package model

type ProductData struct {
	ID          string
	Quantity    int
	ProductName string
	Price       string
	Total       string
}

type SuccessfullyOrderPlayedData struct {
	OrderID  string
	Products []ProductData
}
