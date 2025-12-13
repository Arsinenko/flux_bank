package transaction_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/transaction"
)

type TransactionRepository struct {
	client pb.TransactionServiceClient
}

func (r TransactionRepository) GetRevenue(ctx context.Context, accountId int32, request transaction.GetByDateRange) ([]*transaction.Transaction, error) {
	// req := pb.GetAccountRevenueRequest{
	// 	TargetAccount: accountId,
	// 	DateRange: &pb.GetByDateRangeRequest{
	// 		From: timestamppb.New(request.From),
	// 		To:   timestamppb.New(request.To)},
	// 	PageN:    &request.PageN,
	// 	PageSize: &request.PageSize,
	// }
	// resp, err := r.client.GetAccountRevenue(ctx, &req)
	// if err != nil {
	// 	return nil, fmt.Errorf("transaction_repo.GetAccountRevenue: %w", err)
	// }
	// result := make([]*transaction.Transaction, 0, len(resp.Transactions))
	// for _, t := range resp.Transactions {
	// 	result = append(result, ToTransactionDomain(t))
	// }
	// return result, nil
	panic("not implemented") //TODO implement
}

func NewTransactionRepository(client pb.TransactionServiceClient) TransactionRepository {
	return TransactionRepository{
		client: client,
	}
}

func (r TransactionRepository) GetAll(ctx context.Context) ([]*transaction.Transaction, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("transaction_repo.GetAll: %w", err)
	}
	result := make([]*transaction.Transaction, 0, len(resp.Transactions))
	for _, t := range resp.Transactions {
		result = append(result, ToTransactionDomain(t))
	}
	return result, nil
}

func (r TransactionRepository) GetById(ctx context.Context, id int32) (*transaction.Transaction, error) {
	resp, err := r.client.GetById(ctx, &pb.GetTransactionByIdRequest{TransactionId: id})
	if err != nil {
		return nil, fmt.Errorf("transaction_repo.GetById: %w", err)
	}
	return ToTransactionDomain(resp), nil
}

func (r TransactionRepository) Add(ctx context.Context, transaction *transaction.Transaction) (*transaction.Transaction, error) {
	req := &pb.AddTransactionRequest{
		SourceAccount: transaction.SourceAccount,
		TargetAccount: transaction.TargetAccount,
		Amount:        transaction.Amount,
		Currency:      transaction.Currency,
		Status:        transaction.Status,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("transaction_repo.Add: %w", err)
	}
	return ToTransactionDomain(resp), nil
}

func (r TransactionRepository) Update(ctx context.Context, transaction *transaction.Transaction) error {
	req := &pb.UpdateTransactionRequest{
		TransactionId: transaction.TransactionID,
		SourceAccount: transaction.SourceAccount,
		TargetAccount: transaction.TargetAccount,
		Amount:        transaction.Amount,
		Currency:      transaction.Currency,
		Status:        transaction.Status,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("transaction_repo.Update: %w", err)
	}
	return nil
}

func (r TransactionRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteTransactionRequest{TransactionId: id})
	if err != nil {
		return fmt.Errorf("transaction_repo.Delete: %w", err)
	}
	return nil
}

func (r TransactionRepository) AddBulk(ctx context.Context, transactions []*transaction.Transaction) error {
	var models []*pb.AddTransactionRequest
	for _, t := range transactions {
		models = append(models, &pb.AddTransactionRequest{
			SourceAccount: t.SourceAccount,
			TargetAccount: t.TargetAccount,
			Amount:        t.Amount,
			Currency:      t.Currency,
			Status:        t.Status,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddTransactionBulkRequest{Transactions: models})
	if err != nil {
		return fmt.Errorf("transaction_repo.AddBulk: %w", err)
	}
	return nil
}

func (r TransactionRepository) UpdateBulk(ctx context.Context, transactions []*transaction.Transaction) error {
	var models []*pb.UpdateTransactionRequest
	for _, t := range transactions {
		models = append(models, &pb.UpdateTransactionRequest{
			TransactionId: t.TransactionID,
			SourceAccount: t.SourceAccount,
			TargetAccount: t.TargetAccount,
			Amount:        t.Amount,
			Currency:      t.Currency,
			Status:        t.Status,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateTransactionBulkRequest{Transactions: models})
	if err != nil {
		return fmt.Errorf("transaction_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r TransactionRepository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteTransactionRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteTransactionRequest{TransactionId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteTransactionBulkRequest{Transactions: models})
	if err != nil {
		return fmt.Errorf("transaction_repo.DeleteBulk: %w", err)
	}
	return nil
}
