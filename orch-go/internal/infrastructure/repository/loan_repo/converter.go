package loan_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/loan"
	"time"
)

func ToLoanDomain(p *pb.LoanModel) *loan.Loan {
	if p == nil {
		return nil
	}
	var startDate time.Time
	if p.StartDate != nil {
		startDate = time.Date(int(p.StartDate.Year), time.Month(p.StartDate.Month), int(p.StartDate.Day), 0, 0, 0, 0, time.UTC)
	}
	var endDate time.Time
	if p.EndDate != nil {
		endDate = time.Date(int(p.EndDate.Year), time.Month(p.EndDate.Month), int(p.EndDate.Day), 0, 0, 0, 0, time.UTC)
	}
	return &loan.Loan{
		LoanID:       p.LoanId,
		CustomerID:   p.CustomerId,
		Principal:    p.Principal,
		InterestRate: p.InterestRate,
		StartDate:    &startDate,
		EndDate:      &endDate,
		Status:       p.Status,
	}
}

func FromLoanDomain(l *loan.Loan) *pb.LoanModel {
	if l == nil {
		return nil
	}
	return &pb.LoanModel{
		LoanId:       l.LoanID,
		CustomerId:   l.CustomerID,
		Principal:    l.Principal,
		InterestRate: l.InterestRate,
		Status:       l.Status,
	}
}

func ToLoanPaymentDomain(p *pb.LoanPaymentModel) *loan.LoanPayment {
	if p == nil {
		return nil
	}
	var paymentDate time.Time
	if p.PaymentDate != nil {
		paymentDate = time.Date(int(p.PaymentDate.Year), time.Month(p.PaymentDate.Month), int(p.PaymentDate.Day), 0, 0, 0, 0, time.UTC)
	}
	return &loan.LoanPayment{
		PaymentID:   p.PaymentId,
		LoanID:      p.LoanId,
		Amount:      p.Amount,
		PaymentDate: &paymentDate,
		IsPaid:      p.IsPaid,
	}
}

func FromLoanPaymentDomain(lp *loan.LoanPayment) *pb.LoanPaymentModel {
	if lp == nil {
		return nil
	}
	return &pb.LoanPaymentModel{
		PaymentId: lp.PaymentID,
		LoanId:    lp.LoanID,
		Amount:    lp.Amount,
		IsPaid:    lp.IsPaid,
	}
}

func ToDateOnly(t *time.Time) *pb.DateOnly {
	if t == nil {
		return nil
	}
	return &pb.DateOnly{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}
