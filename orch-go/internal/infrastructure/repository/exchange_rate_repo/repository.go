package exchange_rate_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/exchange_rate"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.ExchangeRateServiceClient
}

func (r Repository) GetByBaseCurrency(ctx context.Context, baseCurrency string) ([]*exchange_rate.ExchangeRate, error) {
	resp, err := r.client.GetByBaseCurrency(ctx, &pb.GetExchangeRateByBaseCurrencyRequest{
		BaseCurrency: baseCurrency,
	})
	if err != nil {
		return nil, fmt.Errorf("exchange_rate_repo.GetByBaseCurrency: %w", err)
	}
	result := make([]*exchange_rate.ExchangeRate, 0, len(resp.ExchangeRates))
	for _, er := range resp.ExchangeRates {
		result = append(result, ToDomain(er))
	}
	return result, nil
}

func NewRepository(client pb.ExchangeRateServiceClient) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*exchange_rate.ExchangeRate, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, fmt.Errorf("exchange_rate_repo.GetAll: %w", err)
	}
	result := make([]*exchange_rate.ExchangeRate, 0, len(resp.ExchangeRates))
	for _, er := range resp.ExchangeRates {
		result = append(result, ToDomain(er))
	}
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*exchange_rate.ExchangeRate, error) {
	resp, err := r.client.GetById(ctx, &pb.GetExchangeRateByIdRequest{RateId: id})
	if err != nil {
		return nil, fmt.Errorf("exchange_rate_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Add(ctx context.Context, exchangeRate *exchange_rate.ExchangeRate) (*exchange_rate.ExchangeRate, error) {
	req := &pb.AddExchangeRateRequest{
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("exchange_rate_repo.Add: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, exchangeRate *exchange_rate.ExchangeRate) error {
	req := &pb.UpdateExchangeRateRequest{
		RateId:         exchangeRate.RateID,
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("exchange_rate_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteExchangeRateRequest{RateId: id})
	if err != nil {
		return fmt.Errorf("exchange_rate_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, exchangeRates []*exchange_rate.ExchangeRate) error {
	var models []*pb.AddExchangeRateRequest
	for _, er := range exchangeRates {
		models = append(models, &pb.AddExchangeRateRequest{
			BaseCurrency:   er.BaseCurrency,
			TargetCurrency: er.TargetCurrency,
			Rate:           er.Rate,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddExchangeRateBulkRequest{ExchangeRates: models})
	if err != nil {
		return fmt.Errorf("exchange_rate_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, exchangeRates []*exchange_rate.ExchangeRate) error {
	var models []*pb.UpdateExchangeRateRequest
	for _, er := range exchangeRates {
		models = append(models, &pb.UpdateExchangeRateRequest{
			RateId:         er.RateID,
			BaseCurrency:   er.BaseCurrency,
			TargetCurrency: er.TargetCurrency,
			Rate:           er.Rate,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateExchangeRateBulkRequest{ExchangeRates: models})
	if err != nil {
		return fmt.Errorf("exchange_rate_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteExchangeRateRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteExchangeRateRequest{RateId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteExchangeRateBulkRequest{ExchangeRates: models})
	if err != nil {
		return fmt.Errorf("exchange_rate_repo.DeleteBulk: %w", err)
	}
	return nil
}
