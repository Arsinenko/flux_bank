package exchange_rate_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/exchange_rate"
	"time"
)

func ToDomain(p *pb.ExchangeRateModel) *exchange_rate.ExchangeRate {
	if p == nil {
		return nil
	}
	var updatedAt time.Time
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt.AsTime()
	}
	return &exchange_rate.ExchangeRate{
		RateID:         p.RateId,
		BaseCurrency:   p.BaseCurrency,
		TargetCurrency: p.TargetCurrency,
		Rate:           p.Rate,
		UpdatedAt:      &updatedAt,
	}
}

func FromDomain(e *exchange_rate.ExchangeRate) *pb.ExchangeRateModel {
	if e == nil {
		return nil
	}
	return &pb.ExchangeRateModel{
		RateId:         e.RateID,
		BaseCurrency:   e.BaseCurrency,
		TargetCurrency: e.TargetCurrency,
		Rate:           e.Rate,
	}
}
