package loan

import "time"

type Loan struct {
	LoanID       int32
	CustomerID   *int32
	Principal    *string
	InterestRate *string
	StartDate    *time.Time
	EndDate      *time.Time
	Status       *string
}

type LoanPayment struct {
	PaymentID   int32
	LoanID      *int32
	Amount      *string
	PaymentDate *time.Time
	IsPaid      *bool
}
