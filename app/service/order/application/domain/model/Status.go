package model

type Status string

const (
	PROCESSING Status = "PROCESSING"
	CONFIRMED  Status = "CONFIRMED"
	CANCELLED  Status = "CANCELLED"
	DELIVERING Status = "DELIVERING"
	COMPLETED  Status = "COMPLETED"
)
