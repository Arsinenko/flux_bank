package loan

import "context"

type LoanRepository interface {
	GetAll(ctx context.Context) ([]*Loan, error)
	GetById(ctx context.Context, id int32) (*Loan, error)
	GetByCustomer(ctx context.Context, customerId int32) ([]*Loan, error)
	Add(ctx context.Context, loan *Loan) (*Loan, error)
	Update(ctx context.Context, loan *Loan) error
	Delete(ctx context.Context, id int32) error
	DeleteBulk(ctx context.Context, ids []int32) error
}

type LoanPaymentRepository interface {
	GetAll(ctx context.Context) ([]*LoanPayment, error)
	GetById(ctx context.Context, id int32) (*LoanPayment, error)
	Add(ctx context.Context, loanPayment *LoanPayment) (*LoanPayment, error)
	Update(ctx context.Context, loanPayment *LoanPayment) error
	Delete(ctx context.Context, id int32) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
