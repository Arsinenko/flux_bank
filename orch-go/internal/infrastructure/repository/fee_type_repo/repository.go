package fee_type_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/fee_type"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.FeeTypeServiceClient
}

func NewRepository(client pb.FeeTypeServiceClient) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*fee_type.FeeType, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, fmt.Errorf("fee_type_repo.GetAll: %w", err)
	}
	result := make([]*fee_type.FeeType, 0, len(resp.FeeTypes))
	for _, ft := range resp.FeeTypes {
		result = append(result, ToDomain(ft))
	}
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*fee_type.FeeType, error) {
	resp, err := r.client.GetById(ctx, &pb.GetFeeTypeByIdRequest{FeeId: id})
	if err != nil {
		return nil, fmt.Errorf("fee_type_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Add(ctx context.Context, feeType *fee_type.FeeType) (*fee_type.FeeType, error) {
	req := &pb.AddFeeTypeRequest{
		Name:        feeType.Name,
		Description: feeType.Description,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("fee_type_repo.Add: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, feeType *fee_type.FeeType) error {
	req := &pb.UpdateFeeTypeRequest{
		FeeId:       feeType.FeeID,
		Name:        feeType.Name,
		Description: feeType.Description,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("fee_type_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteFeeTypeRequest{FeeId: id})
	if err != nil {
		return fmt.Errorf("fee_type_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, feeTypes []*fee_type.FeeType) error {
	var models []*pb.AddFeeTypeRequest
	for _, ft := range feeTypes {
		models = append(models, &pb.AddFeeTypeRequest{
			Name:        ft.Name,
			Description: ft.Description,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddFeeTypeBulkRequest{FeeTypes: models})
	if err != nil {
		return fmt.Errorf("fee_type_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, feeTypes []*fee_type.FeeType) error {
	var models []*pb.UpdateFeeTypeRequest
	for _, ft := range feeTypes {
		models = append(models, &pb.UpdateFeeTypeRequest{
			FeeId:       ft.FeeID,
			Name:        ft.Name,
			Description: ft.Description,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateFeeTypeBulkRequest{FeeTypes: models})
	if err != nil {
		return fmt.Errorf("fee_type_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteFeeTypeRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteFeeTypeRequest{FeeId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteFeeTypeBulkRequest{FeeTypes: models})
	if err != nil {
		return fmt.Errorf("fee_type_repo.DeleteBulk: %w", err)
	}
	return nil
}
