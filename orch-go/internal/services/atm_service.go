package services

import (
	"context"
	"orch-go/internal/domain/atm"
	"orch-go/internal/infrastructure/repository/atm_repo"
)

type AtmService struct {
	repo atm_repo.Repository
}

func NewAtmService(repo atm_repo.Repository) *AtmService {
	return &AtmService{repo: repo}
}

func (s *AtmService) GetAtmById(ctx context.Context, id int32) (*atm.Atm, error) {
	return s.repo.GetById(ctx, id)
}

func (s *AtmService) GetAtmsByStatus(ctx context.Context, status string) ([]*atm.Atm, error) {
	return s.repo.GetByStatus(ctx, status)
}

func (s *AtmService) GetAtmsByLocationSubStr(ctx context.Context, subStr string) ([]*atm.Atm, error) {
	return s.repo.GetByLocationSubStr(ctx, subStr)
}

func (s *AtmService) GetAtmsByBranch(ctx context.Context, branchId int32) ([]*atm.Atm, error) {
	return s.repo.GetByBranch(ctx, branchId)
}

func (s *AtmService) GetAllAtms(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*atm.Atm, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *AtmService) CreateAtm(ctx context.Context, atm *atm.Atm) (*atm.Atm, error) {
	return s.repo.Add(ctx, atm)
}

func (s *AtmService) UpdateAtm(ctx context.Context, atm *atm.Atm) error {
	return s.repo.Update(ctx, atm)
}

func (s *AtmService) DeleteAtm(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *AtmService) CreateAtmBulk(ctx context.Context, atms []*atm.Atm) error {
	return s.repo.AddBulk(ctx, atms)
}

func (s *AtmService) UpdateAtmBulk(ctx context.Context, atms []*atm.Atm) error {
	return s.repo.UpdateBulk(ctx, atms)
}

func (s *AtmService) DeleteAtmBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
