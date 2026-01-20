package services

import (
	"context"
	"orch-go/internal/domain/fee_type"
	"orch-go/internal/infrastructure/repository/fee_type_repo"
)

type FeeTypeService struct {
	repo fee_type_repo.Repository
}

func NewFeeTypeService(repo fee_type_repo.Repository) *FeeTypeService {
	return &FeeTypeService{repo: repo}
}

func (s *FeeTypeService) GetFeeTypeById(ctx context.Context, id int32) (*fee_type.FeeType, error) {
	return s.repo.GetById(ctx, id)
}

func (s *FeeTypeService) GetAllFeeTypes(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*fee_type.FeeType, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *FeeTypeService) CreateFeeType(ctx context.Context, ft *fee_type.FeeType) (*fee_type.FeeType, error) {
	return s.repo.Add(ctx, ft)
}

func (s *FeeTypeService) UpdateFeeType(ctx context.Context, ft *fee_type.FeeType) error {
	return s.repo.Update(ctx, ft)
}

func (s *FeeTypeService) DeleteFeeType(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *FeeTypeService) CreateFeeTypeBulk(ctx context.Context, fts []*fee_type.FeeType) error {
	return s.repo.AddBulk(ctx, fts)
}

func (s *FeeTypeService) UpdateFeeTypeBulk(ctx context.Context, fts []*fee_type.FeeType) error {
	return s.repo.UpdateBulk(ctx, fts)
}

func (s *FeeTypeService) DeleteFeeTypeBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
