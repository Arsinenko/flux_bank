package loan_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/loan"
)

type LoanRepository struct {
	client pb.LoanServiceClient
}

func (r LoanRepository) GetByCustomer(ctx context.Context, customerId int32) ([]*loan.Loan, error) {
	resp, err := r.client.GetByCustomer(ctx, &pb.GetLoansByCustomerRequest{CustomerId: customerId})
	if err != nil {
		return nil, fmt.Errorf("loan_repo.GetByCustomer: %w", err)
	}
	result := make([]*loan.Loan, 0, len(resp.Loans))
	for _, l := range resp.Loans {
		result = append(result, ToLoanDomain(l))
	}
	return result, nil
}

func NewLoanRepository(client pb.LoanServiceClient) LoanRepository {
	return LoanRepository{
		client: client,
	}
}

func (r LoanRepository) GetAll(ctx context.Context, pageN, pageSize int32) ([]*loan.Loan, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("loan_repo.GetAll: %w", err)
	}
	result := make([]*loan.Loan, 0, len(resp.Loans))
	for _, l := range resp.Loans {
		result = append(result, ToLoanDomain(l))
	}
	return result, nil
}

func (r LoanRepository) GetById(ctx context.Context, id int32) (*loan.Loan, error) {
	resp, err := r.client.GetById(ctx, &pb.GetLoanByIdRequest{LoanId: id})
	if err != nil {
		return nil, fmt.Errorf("loan_repo.GetById: %w", err)
	}
	return ToLoanDomain(resp), nil
}

func (r LoanRepository) Add(ctx context.Context, loan *loan.Loan) (*loan.Loan, error) {
	req := &pb.AddLoanRequest{
		CustomerId:   loan.CustomerID,
		Principal:    loan.Principal,
		InterestRate: loan.InterestRate,
		StartDate:    ToDateOnly(loan.StartDate),
		EndDate:      ToDateOnly(loan.EndDate),
		Status:       loan.Status,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("loan_repo.Add: %w", err)
	}
	return ToLoanDomain(resp), nil
}

func (r LoanRepository) Update(ctx context.Context, loan *loan.Loan) error {
	req := &pb.UpdateLoanRequest{
		LoanId:       loan.LoanID,
		CustomerId:   loan.CustomerID,
		Principal:    loan.Principal,
		InterestRate: loan.InterestRate,
		StartDate:    ToDateOnly(loan.StartDate),
		EndDate:      ToDateOnly(loan.EndDate),
		Status:       loan.Status,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("loan_repo.Update: %w", err)
	}
	return nil
}

func (r LoanRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteLoanRequest{LoanId: id})
	if err != nil {
		return fmt.Errorf("loan_repo.Delete: %w", err)
	}
	return nil
}

func (r LoanRepository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteLoanRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteLoanRequest{LoanId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteLoanBulkRequest{Loans: models})
	if err != nil {
		return fmt.Errorf("loan_repo.DeleteBulk: %w", err)
	}
	return nil
}
