package atm_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/atm"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.AtmServiceClient
}

func (r Repository) GetByStatus(ctx context.Context, status string) ([]*atm.Atm, error) {
	resp, err := r.client.GetByStatus(ctx, &pb.GetAtmsByStatusRequest{Status: status})
	if err != nil {
		return nil, fmt.Errorf("atm_repo.GetByStatus: %w", err)
	}
	result := ToAtmsDomain(resp)
	return result, nil

}

func (r Repository) GetByLocationSubStr(ctx context.Context, subStr string) ([]*atm.Atm, error) {
	resp, err := r.client.GetByLocationSubStr(ctx, &pb.GetAtmsByLocationSubStrRequest{SubStr: subStr})
	if err != nil {
		return nil, fmt.Errorf("atm_repo.GetByLocationSubStr: %w", err)
	}
	result := ToAtmsDomain(resp)
	return result, nil

}

func (r Repository) GetByBranch(ctx context.Context, branchId int32) ([]*atm.Atm, error) {
	resp, err := r.client.GetByBranch(ctx, &pb.GetAtmsByBranchRequest{BranchId: branchId})
	if err != nil {
		return nil, fmt.Errorf("atm_repo.GetByBranch: %w", err)
	}
	result := ToAtmsDomain(resp)
	return result, nil
}

func NewRepository(client pb.AtmServiceClient) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*atm.Atm, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, fmt.Errorf("atm_repo.GetAll: %w", err)
	}
	result := ToAtmsDomain(resp)
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*atm.Atm, error) {
	resp, err := r.client.GetById(ctx, &pb.GetAtmByIdRequest{AtmId: id})
	if err != nil {
		return nil, fmt.Errorf("atm_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Add(ctx context.Context, atm *atm.Atm) (*atm.Atm, error) {
	req := &pb.AddAtmRequest{
		BranchId: atm.BranchID,
		Location: atm.Location,
		Status:   atm.Status,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("atm_repo.Add: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, atm *atm.Atm) error {
	req := &pb.UpdateAtmRequest{
		AtmId:    atm.AtmID,
		BranchId: atm.BranchID,
		Location: atm.Location,
		Status:   atm.Status,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("atm_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteAtmRequest{AtmId: id})
	if err != nil {
		return fmt.Errorf("atm_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, atms []*atm.Atm) error {
	var models []*pb.AddAtmRequest
	for _, a := range atms {
		models = append(models, &pb.AddAtmRequest{
			BranchId: a.BranchID,
			Location: a.Location,
			Status:   a.Status,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddAtmBulkRequest{Atms: models})
	if err != nil {
		return fmt.Errorf("atm_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, atms []*atm.Atm) error {
	var models []*pb.UpdateAtmRequest
	for _, a := range atms {
		models = append(models, &pb.UpdateAtmRequest{
			AtmId:    a.AtmID,
			BranchId: a.BranchID,
			Location: a.Location,
			Status:   a.Status,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateAtmBulkRequest{Atms: models})
	if err != nil {
		return fmt.Errorf("atm_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteAtmRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteAtmRequest{AtmId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteAtmBulkRequest{Atms: models})
	if err != nil {
		return fmt.Errorf("atm_repo.DeleteBulk: %w", err)
	}
	return nil
}
