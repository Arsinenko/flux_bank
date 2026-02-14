package transaction

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	TransactionID int32
	SourceAccount *int32
	TargetAccount *int32
	Amount        decimal.Decimal
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
	Amount        *decimal.Decimal
}

type GetByDateRange struct {
	From, To        time.Time
	PageN, PageSize int32
}
