package notification

import "time"

type Notification struct {
	Id         int32
	CustomerId int32
	Message    string
	CreatedAt  time.Time
	IsRead     bool
}
