package in

type ConfirmOrderPort interface {
	ConfirmOrder(orderId string) error
}
