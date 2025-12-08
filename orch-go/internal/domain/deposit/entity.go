package deposit

import "time"

type Deposit struct {
	DepositID    int32
	CustomerID   int32
	Amount       string
	InterestRate string
	StartDate    time.Time
	EndDate      time.Time
	Status       string
}
