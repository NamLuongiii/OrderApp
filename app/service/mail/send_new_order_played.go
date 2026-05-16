package mail

func (m *ServiceImpl) SendNewOrderPlayed(to string, command SendNewOrderPlayedCommand) {
	m.sendGMail(
		to,
		"Order Updated",
		"template/send_new_order_played.html",
		command)
}

type ProductData struct {
	ID          string
	Quantity    int64
	ProductName string
	Price       int64
	Total       int64
}

type SendNewOrderPlayedCommand struct {
	OrderID  string
	Products []ProductData
}
