package services

import (
	"context"
	"orch-go/internal/domain/branch"
	"orch-go/internal/infrastructure/repository/branch_repo"
)

type BranchService struct {
	repo branch_repo.Repository
}

func NewBranchService(repo branch_repo.Repository) *BranchService {
	return &BranchService{repo: repo}
}

func (s *BranchService) GetBranchById(ctx context.Context, id int32) (*branch.Branch, error) {
	return s.repo.GetById(ctx, id)
}

func (s *BranchService) GetAllBranches(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*branch.Branch, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *BranchService) CreateBranch(ctx context.Context, branch *branch.Branch) (*branch.Branch, error) {
	return s.repo.Add(ctx, branch)
}

func (s *BranchService) UpdateBranch(ctx context.Context, branch *branch.Branch) error {
	return s.repo.Update(ctx, branch)
}

func (s *BranchService) DeleteBranch(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *BranchService) CreateBranchBulk(ctx context.Context, branches []*branch.Branch) error {
	return s.repo.AddBulk(ctx, branches)
}

func (s *BranchService) UpdateBranchBulk(ctx context.Context, branches []*branch.Branch) error {
	return s.repo.UpdateBulk(ctx, branches)
}

func (s *BranchService) DeleteBranchBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
