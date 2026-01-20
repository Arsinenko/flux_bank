package services

import (
	"context"
	"orch-go/internal/domain/deposit"
	"orch-go/internal/infrastructure/repository/deposit_repo"
)

type DepositService struct {
	repo deposit_repo.Repository
}

func NewDepositService(repo deposit_repo.Repository) *DepositService {
	return &DepositService{repo: repo}
}

func (s *DepositService) GetDepositById(ctx context.Context, id int32) (*deposit.Deposit, error) {
	return s.repo.GetById(ctx, id)
}

func (s *DepositService) GetDepositsByCustomer(ctx context.Context, customerId int32) ([]*deposit.Deposit, error) {
	return s.repo.GetByCustomer(ctx, customerId)
}

func (s *DepositService) GetAllDeposits(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*deposit.Deposit, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *DepositService) CreateDeposit(ctx context.Context, deposit *deposit.Deposit) (*deposit.Deposit, error) {
	return s.repo.Add(ctx, deposit)
}

func (s *DepositService) UpdateDeposit(ctx context.Context, deposit *deposit.Deposit) error {
	return s.repo.Update(ctx, deposit)
}

func (s *DepositService) DeleteDeposit(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *DepositService) CreateDepositBulk(ctx context.Context, deposits []*deposit.Deposit) error {
	return s.repo.AddBulk(ctx, deposits)
}

func (s *DepositService) UpdateDepositBulk(ctx context.Context, deposits []*deposit.Deposit) error {
	return s.repo.UpdateBulk(ctx, deposits)
}

func (s *DepositService) DeleteDepositBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
