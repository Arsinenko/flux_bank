package services

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"orch-go/internal/domain/transaction"
	"orch-go/internal/infrastructure/repository/account/account_repo"
	"orch-go/internal/infrastructure/repository/transaction_repo"
)

type TransactionService struct {
	txRepo       transaction_repo.TransactionRepository
	categoryRepo transaction_repo.TransactionCategoryRepository
	feeRepo      transaction_repo.TransactionFeeRepository
	accRepo      account_repo.Repository
}

func NewTransactionService(
	txRepo transaction_repo.TransactionRepository,
	categoryRepo transaction_repo.TransactionCategoryRepository,
	feeRepo transaction_repo.TransactionFeeRepository,
	accRepo account_repo.Repository,
) *TransactionService {
	return &TransactionService{
		txRepo:       txRepo,
		categoryRepo: categoryRepo,
		feeRepo:      feeRepo,
		accRepo:      accRepo,
	}
}

// Transaction methods

func (s *TransactionService) GetTransactionRevenue(ctx context.Context, accountId int32, req transaction.GetByDateRange) ([]*transaction.Transaction, error) {
	return s.txRepo.GetRevenue(ctx, accountId, req)
}

func (s *TransactionService) GetAllTransactions(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.Transaction, error) {
	return s.txRepo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *TransactionService) GetTransactionById(ctx context.Context, id int32) (*transaction.Transaction, error) {
	return s.txRepo.GetById(ctx, id)
}

func (s *TransactionService) CreateTransaction(ctx context.Context, t *transaction.Transaction) (*transaction.Transaction, error) {
	sourceAcc, err := s.accRepo.GetById(ctx, *t.SourceAccount)
	if err != nil {
		return nil, err
	}
	if sourceAcc.Balance.LessThan(t.Amount) {
		return nil, fmt.Errorf("insufficient balance")
	}
	targetAcc, err := s.accRepo.GetById(ctx, *t.TargetAccount)
	if err != nil {
		return nil, err
	}
	sourceAcc.Balance = sourceAcc.Balance.Sub(t.Amount)
	targetAcc.Balance = targetAcc.Balance.Add(t.Amount)

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.accRepo.Update(gCtx, sourceAcc)
	})

	g.Go(func() error {
		return s.accRepo.Update(gCtx, targetAcc)
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return s.txRepo.Add(ctx, t)

}

func (s *TransactionService) UpdateTransaction(ctx context.Context, t *transaction.Transaction) error {
	return s.txRepo.Update(ctx, t)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id int32) error {
	return s.txRepo.Delete(ctx, id)
}

func (s *TransactionService) CreateTransactionBulk(ctx context.Context, ts []*transaction.Transaction) error {
	return s.txRepo.AddBulk(ctx, ts)
}

func (s *TransactionService) UpdateTransactionBulk(ctx context.Context, ts []*transaction.Transaction) error {
	return s.txRepo.UpdateBulk(ctx, ts)
}

func (s *TransactionService) DeleteTransactionBulk(ctx context.Context, ids []int32) error {
	return s.txRepo.DeleteBulk(ctx, ids)
}

// Transaction Category methods

func (s *TransactionService) GetAllTransactionCategories(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.TransactionCategory, error) {
	return s.categoryRepo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *TransactionService) GetTransactionCategoryById(ctx context.Context, id int32) (*transaction.TransactionCategory, error) {
	return s.categoryRepo.GetById(ctx, id)
}

func (s *TransactionService) CreateTransactionCategory(ctx context.Context, tc *transaction.TransactionCategory) (*transaction.TransactionCategory, error) {
	return s.categoryRepo.Add(ctx, tc)
}

func (s *TransactionService) UpdateTransactionCategory(ctx context.Context, tc *transaction.TransactionCategory) error {
	return s.categoryRepo.Update(ctx, tc)
}

func (s *TransactionService) DeleteTransactionCategory(ctx context.Context, id int32) error {
	return s.categoryRepo.Delete(ctx, id)
}

func (s *TransactionService) CreateTransactionCategoryBulk(ctx context.Context, tcs []*transaction.TransactionCategory) error {
	return s.categoryRepo.AddBulk(ctx, tcs)
}

func (s *TransactionService) UpdateTransactionCategoryBulk(ctx context.Context, tcs []*transaction.TransactionCategory) error {
	return s.categoryRepo.UpdateBulk(ctx, tcs)
}

func (s *TransactionService) DeleteTransactionCategoryBulk(ctx context.Context, ids []int32) error {
	return s.categoryRepo.DeleteBulk(ctx, ids)
}

// Transaction Fee methods

func (s *TransactionService) GetAllTransactionFees(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.TransactionFee, error) {
	return s.feeRepo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *TransactionService) GetTransactionFeeById(ctx context.Context, id int32) (*transaction.TransactionFee, error) {
	return s.feeRepo.GetById(ctx, id)
}

func (s *TransactionService) CreateTransactionFee(ctx context.Context, tf *transaction.TransactionFee) (*transaction.TransactionFee, error) {
	return s.feeRepo.Add(ctx, tf)
}

func (s *TransactionService) UpdateTransactionFee(ctx context.Context, tf *transaction.TransactionFee) error {
	return s.feeRepo.Update(ctx, tf)
}

func (s *TransactionService) DeleteTransactionFee(ctx context.Context, id int32) error {
	return s.feeRepo.Delete(ctx, id)
}

func (s *TransactionService) CreateTransactionFeeBulk(ctx context.Context, tfs []*transaction.TransactionFee) error {
	return s.feeRepo.AddBulk(ctx, tfs)
}

func (s *TransactionService) UpdateTransactionFeeBulk(ctx context.Context, tfs []*transaction.TransactionFee) error {
	return s.feeRepo.UpdateBulk(ctx, tfs)
}

func (s *TransactionService) DeleteTransactionFeeBulk(ctx context.Context, ids []int32) error {
	return s.feeRepo.DeleteBulk(ctx, ids)
}
