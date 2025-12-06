package branch_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/branch"
)

type Repository struct {
	client pb.BranchServiceClient
}

func NewRepository(client pb.BranchServiceClient) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context) ([]*branch.Branch, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("branch_repo.GetAll: %w", err)
	}
	result := make([]*branch.Branch, 0, len(resp.Branches))
	for _, b := range resp.Branches {
		result = append(result, ToDomain(b))
	}
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*branch.Branch, error) {
	resp, err := r.client.GetById(ctx, &pb.GetBranchByIdRequest{BranchId: id})
	if err != nil {
		return nil, fmt.Errorf("branch_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Add(ctx context.Context, branch *branch.Branch) (*branch.Branch, error) {
	req := &pb.AddBranchRequest{
		Name:    branch.Name,
		City:    branch.City,
		Address: branch.Address,
		Phone:   branch.Phone,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("branch_repo.Add: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, branch *branch.Branch) error {
	req := &pb.UpdateBranchRequest{
		BranchId: branch.BranchID,
		Name:     branch.Name,
		City:     branch.City,
		Address:  branch.Address,
		Phone:    branch.Phone,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("branch_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteBranchRequest{BranchId: id})
	if err != nil {
		return fmt.Errorf("branch_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, branches []*branch.Branch) error {
	var models []*pb.AddBranchRequest
	for _, b := range branches {
		models = append(models, &pb.AddBranchRequest{
			Name:    b.Name,
			City:    b.City,
			Address: b.Address,
			Phone:   b.Phone,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddBranchBulkRequest{Branches: models})
	if err != nil {
		return fmt.Errorf("branch_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, branches []*branch.Branch) error {
	var models []*pb.UpdateBranchRequest
	for _, b := range branches {
		models = append(models, &pb.UpdateBranchRequest{
			BranchId: b.BranchID,
			Name:     b.Name,
			City:     b.City,
			Address:  b.Address,
			Phone:    b.Phone,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateBranchBulkRequest{Branches: models})
	if err != nil {
		return fmt.Errorf("branch_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteBranchRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteBranchRequest{BranchId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteBranchBulkRequest{Branches: models})
	if err != nil {
		return fmt.Errorf("branch_repo.DeleteBulk: %w", err)
	}
	return nil
}
