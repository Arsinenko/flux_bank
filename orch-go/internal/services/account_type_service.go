package services

import (
	"context"
	"orch-go/internal/domain/account"
	"orch-go/internal/infrastructure/repository/account/account_repo"
)

type AccountTypeService struct {
	repo account_repo.AccountTypeRepository
}

func NewAccountTypeService(repo account_repo.AccountTypeRepository) *AccountTypeService {
	return &AccountTypeService{repo: repo}
}

func (s *AccountTypeService) GetAccountTypeById(ctx context.Context, id int32) (*account.AccountType, error) {
	return s.repo.GetById(ctx, id)
}

func (s *AccountTypeService) GetAllAccountTypes(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]account.AccountType, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *AccountTypeService) CreateAccountType(ctx context.Context, at *account.AccountType) (*account.AccountType, error) {
	return s.repo.Create(ctx, at)
}

func (s *AccountTypeService) UpdateAccountType(ctx context.Context, at *account.AccountType) error {
	return s.repo.Update(ctx, at)
}

func (s *AccountTypeService) DeleteAccountType(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *AccountTypeService) CreateAccountTypeBulk(ctx context.Context, ats []account.AccountType) error {
	return s.repo.AddBulk(ctx, ats)
}

func (s *AccountTypeService) UpdateAccountTypeBulk(ctx context.Context, ats []account.AccountType) error {
	return s.repo.UpdateBulk(ctx, ats)
}

func (s *AccountTypeService) DeleteAccountTypeBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
