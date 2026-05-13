package in

type MarkOrderCompletedPort interface {
	MarkOrderCompleted(orderId string) error
}
