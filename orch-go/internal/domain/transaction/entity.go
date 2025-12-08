package transaction

import "time"

type Transaction struct {
	TransactionID int32
	SourceAccount *int32
	TargetAccount *int32
	Amount        string
	Currency      string
	CreatedAt     *time.Time
	Status        *string
}

type TransactionCategory struct {
	CategoryID int32
	Name       string
}

type TransactionFee struct {
	ID            int32
	TransactionID *int32
	FeeID         *int32
	Amount        *string
}

type GetByDateRange struct {
	From, To        time.Time
	PageN, PageSize int32
}
