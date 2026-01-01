package services

import (
	"context"
	"orch-go/internal/domain/exchange_rate"
	"orch-go/internal/infrastructure/repository/exchange_rate_repo"
)

type ExchangeRateService struct {
	repo exchange_rate_repo.Repository
}

func NewExchangeRateService(repo exchange_rate_repo.Repository) *ExchangeRateService {
	return &ExchangeRateService{repo: repo}
}

func (s *ExchangeRateService) GetExchangeRateById(ctx context.Context, id int32) (*exchange_rate.ExchangeRate, error) {
	return s.repo.GetById(ctx, id)
}

func (s *ExchangeRateService) GetExchangeRatesByBaseCurrency(ctx context.Context, baseCurrency string) ([]*exchange_rate.ExchangeRate, error) {
	return s.repo.GetByBaseCurrency(ctx, baseCurrency)
}

func (s *ExchangeRateService) GetAllExchangeRates(ctx context.Context, pageN, pageSize int32) ([]*exchange_rate.ExchangeRate, error) {
	return s.repo.GetAll(ctx, pageN, pageSize)
}

func (s *ExchangeRateService) CreateExchangeRate(ctx context.Context, er *exchange_rate.ExchangeRate) (*exchange_rate.ExchangeRate, error) {
	return s.repo.Add(ctx, er)
}

func (s *ExchangeRateService) UpdateExchangeRate(ctx context.Context, er *exchange_rate.ExchangeRate) error {
	return s.repo.Update(ctx, er)
}

func (s *ExchangeRateService) DeleteExchangeRate(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *ExchangeRateService) CreateExchangeRateBulk(ctx context.Context, ers []*exchange_rate.ExchangeRate) error {
	return s.repo.AddBulk(ctx, ers)
}

func (s *ExchangeRateService) UpdateExchangeRateBulk(ctx context.Context, ers []*exchange_rate.ExchangeRate) error {
	return s.repo.UpdateBulk(ctx, ers)
}

func (s *ExchangeRateService) DeleteExchangeRateBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
