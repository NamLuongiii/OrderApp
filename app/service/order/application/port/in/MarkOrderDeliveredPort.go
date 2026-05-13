package in

type MarkOrderDeliveredPort interface {
	MarkOrderDelivered(orderId string) error
}
