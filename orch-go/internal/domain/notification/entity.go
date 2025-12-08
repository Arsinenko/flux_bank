package notification

import "time"

type Notification struct {
	Id         int32
	CustomerId int32
	Message    string
	CreatedAt  time.Time
	IsRead     bool
}

type GetByDateRangeRequest struct {
	From, To        time.Time
	PageN, PageSize int32
}
