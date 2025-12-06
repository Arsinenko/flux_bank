package loan_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/loan"
)

type LoanPaymentRepository struct {
	client pb.LoanPaymentServiceClient
}

func NewLoanPaymentRepository(client pb.LoanPaymentServiceClient) LoanPaymentRepository {
	return LoanPaymentRepository{
		client: client,
	}
}

func (r LoanPaymentRepository) GetAll(ctx context.Context) ([]*loan.LoanPayment, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("loan_payment_repo.GetAll: %w", err)
	}
	result := make([]*loan.LoanPayment, 0, len(resp.LoanPayments))
	for _, lp := range resp.LoanPayments {
		result = append(result, ToLoanPaymentDomain(lp))
	}
	return result, nil
}

func (r LoanPaymentRepository) GetById(ctx context.Context, id int32) (*loan.LoanPayment, error) {
	resp, err := r.client.GetById(ctx, &pb.GetLoanPaymentByIdRequest{PaymentId: id})
	if err != nil {
		return nil, fmt.Errorf("loan_payment_repo.GetById: %w", err)
	}
	return ToLoanPaymentDomain(resp), nil
}

func (r LoanPaymentRepository) Add(ctx context.Context, loanPayment *loan.LoanPayment) (*loan.LoanPayment, error) {
	req := &pb.AddLoanPaymentRequest{
		LoanId:      loanPayment.LoanID,
		Amount:      loanPayment.Amount,
		PaymentDate: ToDateOnly(loanPayment.PaymentDate),
		IsPaid:      loanPayment.IsPaid,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("loan_payment_repo.Add: %w", err)
	}
	return ToLoanPaymentDomain(resp), nil
}

func (r LoanPaymentRepository) Update(ctx context.Context, loanPayment *loan.LoanPayment) error {
	req := &pb.UpdateLoanPaymentRequest{
		PaymentId:   loanPayment.PaymentID,
		LoanId:      loanPayment.LoanID,
		Amount:      loanPayment.Amount,
		PaymentDate: ToDateOnly(loanPayment.PaymentDate),
		IsPaid:      loanPayment.IsPaid,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("loan_payment_repo.Update: %w", err)
	}
	return nil
}

func (r LoanPaymentRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteLoanPaymentRequest{PaymentId: id})
	if err != nil {
		return fmt.Errorf("loan_payment_repo.Delete: %w", err)
	}
	return nil
}

func (r LoanPaymentRepository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteLoanPaymentRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteLoanPaymentRequest{PaymentId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteLoanPaymentBulkRequest{Payments: models})
	if err != nil {
		return fmt.Errorf("loan_payment_repo.DeleteBulk: %w", err)
	}
	return nil
}
