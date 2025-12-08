package account

import "time"

type Account struct {
	Id         int32
	CustomerId int32
	TypeId     int32
	Iban       string
	Balance    string
	CreatedAt  time.Time
	IsActive   bool
}

type AccountType struct {
	Id          int32
	Name        string
	Description *string
}

type GetByDateRange struct {
	From, To        time.Time
	PageN, PageSize int32
}
