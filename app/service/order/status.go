package order

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusConfirmed Status = "CONFIRMED"
	StatusDelivered Status = "DELIVERED"
	StatusCanceled  Status = "CANCELED"
	StatusCompleted Status = "COMPLETED"
)
