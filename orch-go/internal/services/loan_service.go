package services

import (
	"context"
	"orch-go/internal/domain/loan"
	"orch-go/internal/infrastructure/repository/loan_repo"
)

type LoanService struct {
	loanRepo        loan_repo.LoanRepository
	loanPaymentRepo loan_repo.LoanPaymentRepository
}

func NewLoanService(loanRepo loan_repo.LoanRepository, loanPaymentRepo loan_repo.LoanPaymentRepository) *LoanService {
	return &LoanService{
		loanRepo:        loanRepo,
		loanPaymentRepo: loanPaymentRepo,
	}
}

// Loan methods

func (s *LoanService) GetLoanById(ctx context.Context, id int32) (*loan.Loan, error) {
	return s.loanRepo.GetById(ctx, id)
}

func (s *LoanService) GetLoansByCustomer(ctx context.Context, customerId int32) ([]*loan.Loan, error) {
	return s.loanRepo.GetByCustomer(ctx, customerId)
}

func (s *LoanService) GetAllLoans(ctx context.Context, pageN, pageSize int32) ([]*loan.Loan, error) {
	return s.loanRepo.GetAll(ctx, pageN, pageSize)
}

func (s *LoanService) CreateLoan(ctx context.Context, l *loan.Loan) (*loan.Loan, error) {
	return s.loanRepo.Add(ctx, l)
}

func (s *LoanService) UpdateLoan(ctx context.Context, l *loan.Loan) error {
	return s.loanRepo.Update(ctx, l)
}

func (s *LoanService) DeleteLoan(ctx context.Context, id int32) error {
	return s.loanRepo.Delete(ctx, id)
}

func (s *LoanService) DeleteLoanBulk(ctx context.Context, ids []int32) error {
	return s.loanRepo.DeleteBulk(ctx, ids)
}

// LoanPayment methods

func (s *LoanService) GetLoanPaymentById(ctx context.Context, id int32) (*loan.LoanPayment, error) {
	return s.loanPaymentRepo.GetById(ctx, id)
}

func (s *LoanService) GetLoanPaymentsByLoan(ctx context.Context, loanId int32) ([]*loan.LoanPayment, error) {
	return s.loanPaymentRepo.GetByLoan(ctx, loanId)
}

func (s *LoanService) GetAllLoanPayments(ctx context.Context, pageN, pageSize int32) ([]*loan.LoanPayment, error) {
	return s.loanPaymentRepo.GetAll(ctx, pageN, pageSize)
}

func (s *LoanService) CreateLoanPayment(ctx context.Context, lp *loan.LoanPayment) (*loan.LoanPayment, error) {
	return s.loanPaymentRepo.Add(ctx, lp)
}

func (s *LoanService) UpdateLoanPayment(ctx context.Context, lp *loan.LoanPayment) error {
	return s.loanPaymentRepo.Update(ctx, lp)
}

func (s *LoanService) DeleteLoanPayment(ctx context.Context, id int32) error {
	return s.loanPaymentRepo.Delete(ctx, id)
}

func (s *LoanService) DeleteLoanPaymentBulk(ctx context.Context, ids []int32) error {
	return s.loanPaymentRepo.DeleteBulk(ctx, ids)
}
