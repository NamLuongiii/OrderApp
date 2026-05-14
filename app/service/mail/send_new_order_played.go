package mail

func (m *ServiceImpl) SendNewOrderPlayed(to string, command SendNewOrderPlayedCommand) {
	m.sendGMail(
		to,
		"Order Updated",
		"successfully_order_played",
		command)
}

type ProductData struct {
	ID          string
	Quantity    int
	ProductName string
	Price       string
	Total       string
}

type SendNewOrderPlayedCommand struct {
	OrderID  string
	Products []ProductData
}
