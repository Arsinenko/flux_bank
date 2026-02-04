package services

import (
	"context"
	"orch-go/internal/domain/account"
	"orch-go/internal/infrastructure/repository/account/account_repo"
)

type AccountService struct {
	repo account_repo.Repository
}

func NewAccountService(repo account_repo.Repository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) GetAccountById(ctx context.Context, id int32) (*account.Account, error) {
	return s.repo.GetById(ctx, id)
}

func (s *AccountService) GetAccountsByCustomer(ctx context.Context, customerId int32) ([]*account.Account, error) {
	return s.repo.GetByCustomerId(ctx, customerId)
}

func (s *AccountService) GetAccountsByDateRange(ctx context.Context, req account.GetByDateRange) ([]*account.Account, error) {
	return s.repo.GetByDateRange(ctx, req)
}

func (s *AccountService) CreateAccount(ctx context.Context, acc *account.Account) (*account.Account, error) {
	return s.repo.Create(ctx, acc)
}

func (s *AccountService) UpdateAccount(ctx context.Context, acc *account.Account) error {
	return s.repo.Update(ctx, acc)
}

func (s *AccountService) DeleteAccount(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *AccountService) GetAllAccounts(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]account.Account, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *AccountService) DeleteAccountBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}

func (s *AccountService) CreateAccountBulk(ctx context.Context, accs []*account.Account) error {
	return s.repo.CreateBulk(ctx, accs)
}

func (s *AccountService) UpdateAccountBulk(ctx context.Context, accs []account.Account) error {
	return s.repo.UpdateBulk(ctx, accs)
}
