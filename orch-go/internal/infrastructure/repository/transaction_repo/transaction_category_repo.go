package transaction_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/transaction"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TransactionCategoryRepository struct {
	client pb.TransactionCategoryServiceClient
}

func NewTransactionCategoryRepository(client pb.TransactionCategoryServiceClient) TransactionCategoryRepository {
	return TransactionCategoryRepository{
		client: client,
	}
}

func (r TransactionCategoryRepository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.TransactionCategory, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, fmt.Errorf("transaction_category_repo.GetAll: %w", err)
	}
	result := make([]*transaction.TransactionCategory, 0, len(resp.TransactionCategories))
	for _, tc := range resp.TransactionCategories {
		result = append(result, ToTransactionCategoryDomain(tc))
	}
	return result, nil
}

func (r TransactionCategoryRepository) GetById(ctx context.Context, id int32) (*transaction.TransactionCategory, error) {
	resp, err := r.client.GetById(ctx, &pb.GetTransactionCategoryByIdRequest{CategoryId: id})
	if err != nil {
		return nil, fmt.Errorf("transaction_category_repo.GetById: %w", err)
	}
	return ToTransactionCategoryDomain(resp), nil
}

func (r TransactionCategoryRepository) Add(ctx context.Context, transactionCategory *transaction.TransactionCategory) (*transaction.TransactionCategory, error) {
	req := &pb.AddTransactionCategoryRequest{
		Name: transactionCategory.Name,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("transaction_category_repo.Add: %w", err)
	}
	return ToTransactionCategoryDomain(resp), nil
}

func (r TransactionCategoryRepository) Update(ctx context.Context, transactionCategory *transaction.TransactionCategory) error {
	req := &pb.UpdateTransactionCategoryRequest{
		CategoryId: transactionCategory.CategoryID,
		Name:       transactionCategory.Name,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("transaction_category_repo.Update: %w", err)
	}
	return nil
}

func (r TransactionCategoryRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteTransactionCategoryRequest{CategoryId: id})
	if err != nil {
		return fmt.Errorf("transaction_category_repo.Delete: %w", err)
	}
	return nil
}

func (r TransactionCategoryRepository) AddBulk(ctx context.Context, transactionCategories []*transaction.TransactionCategory) error {
	var models []*pb.AddTransactionCategoryRequest
	for _, tc := range transactionCategories {
		models = append(models, &pb.AddTransactionCategoryRequest{
			Name: tc.Name,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddTransactionCategoryBulkRequest{TransactionCategories: models})
	if err != nil {
		return fmt.Errorf("transaction_category_repo.AddBulk: %w", err)
	}
	return nil
}

func (r TransactionCategoryRepository) UpdateBulk(ctx context.Context, transactionCategories []*transaction.TransactionCategory) error {
	var models []*pb.UpdateTransactionCategoryRequest
	for _, tc := range transactionCategories {
		models = append(models, &pb.UpdateTransactionCategoryRequest{
			CategoryId: tc.CategoryID,
			Name:       tc.Name,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateTransactionCategoryBulkRequest{TransactionCategories: models})
	if err != nil {
		return fmt.Errorf("transaction_category_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r TransactionCategoryRepository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteTransactionCategoryRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteTransactionCategoryRequest{CategoryId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteTransactionCategoryBulkRequest{TransactionCategories: models})
	if err != nil {
		return fmt.Errorf("transaction_category_repo.DeleteBulk: %w", err)
	}
	return nil
}
