package in

type CancelOrderPort interface {
	CancelOrder(orderId string) error
}
