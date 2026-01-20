package transaction_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/transaction"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TransactionFeeRepository struct {
	client pb.TransactionFeeServiceClient
}

func NewTransactionFeeRepository(client pb.TransactionFeeServiceClient) TransactionFeeRepository {
	return TransactionFeeRepository{
		client: client,
	}
}

func (r TransactionFeeRepository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.TransactionFee, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, fmt.Errorf("transaction_fee_repo.GetAll: %w", err)
	}
	result := make([]*transaction.TransactionFee, 0, len(resp.TransactionFees))
	for _, tf := range resp.TransactionFees {
		result = append(result, ToTransactionFeeDomain(tf))
	}
	return result, nil
}

func (r TransactionFeeRepository) GetById(ctx context.Context, id int32) (*transaction.TransactionFee, error) {
	resp, err := r.client.GetById(ctx, &pb.GetTransactionFeeByIdRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("transaction_fee_repo.GetById: %w", err)
	}
	return ToTransactionFeeDomain(resp), nil
}

func (r TransactionFeeRepository) Add(ctx context.Context, transactionFee *transaction.TransactionFee) (*transaction.TransactionFee, error) {
	req := &pb.AddTransactionFeeRequest{
		TransactionId: transactionFee.TransactionID,
		FeeId:         transactionFee.FeeID,
		Amount:        transactionFee.Amount,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("transaction_fee_repo.Add: %w", err)
	}
	return ToTransactionFeeDomain(resp), nil
}

func (r TransactionFeeRepository) Update(ctx context.Context, transactionFee *transaction.TransactionFee) error {
	req := &pb.UpdateTransactionFeeRequest{
		Id:            transactionFee.ID,
		TransactionId: transactionFee.TransactionID,
		FeeId:         transactionFee.FeeID,
		Amount:        transactionFee.Amount,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("transaction_fee_repo.Update: %w", err)
	}
	return nil
}

func (r TransactionFeeRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteTransactionFeeRequest{Id: id})
	if err != nil {
		return fmt.Errorf("transaction_fee_repo.Delete: %w", err)
	}
	return nil
}

func (r TransactionFeeRepository) AddBulk(ctx context.Context, transactionFees []*transaction.TransactionFee) error {
	var models []*pb.AddTransactionFeeRequest
	for _, tf := range transactionFees {
		models = append(models, &pb.AddTransactionFeeRequest{
			TransactionId: tf.TransactionID,
			FeeId:         tf.FeeID,
			Amount:        tf.Amount,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddTransactionFeeBulkRequest{TransactionFees: models})
	if err != nil {
		return fmt.Errorf("transaction_fee_repo.AddBulk: %w", err)
	}
	return nil
}

func (r TransactionFeeRepository) UpdateBulk(ctx context.Context, transactionFees []*transaction.TransactionFee) error {
	var models []*pb.UpdateTransactionFeeRequest
	for _, tf := range transactionFees {
		models = append(models, &pb.UpdateTransactionFeeRequest{
			Id:            tf.ID,
			TransactionId: tf.TransactionID,
			FeeId:         tf.FeeID,
			Amount:        tf.Amount,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateTransactionFeeBulkRequest{TransactionFees: models})
	if err != nil {
		return fmt.Errorf("transaction_fee_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r TransactionFeeRepository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteTransactionFeeRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteTransactionFeeRequest{Id: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteTransactionFeeBulkRequest{TransactionFees: models})
	if err != nil {
		return fmt.Errorf("transaction_fee_repo.DeleteBulk: %w", err)
	}
	return nil
}
