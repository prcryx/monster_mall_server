package types

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Canceled Status = "canceled"
)
