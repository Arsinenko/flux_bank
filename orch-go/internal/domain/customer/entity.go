package customer

import "time"

type Customer struct {
	Id        int32
	FirstName string
	LastName  string
	Email     string
	Phone     *string
	BirthDate *time.Time
	CreatedAt *time.Time
}

type GetBySubStrRequest struct {
	SubStr          string
	PageN, PageSize int32
	Order           string
	Desk            bool
}

type GetByDateRangeRequest struct {
	From, To        time.Time
	PageN, PageSize int32
}
