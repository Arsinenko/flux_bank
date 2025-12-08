package exchange_rate

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*ExchangeRate, error)
	GetById(ctx context.Context, id int32) (*ExchangeRate, error)
	GetByBaseCurrency(ctx context.Context, baseCurrency string) ([]*ExchangeRate, error)
	Add(ctx context.Context, exchangeRate *ExchangeRate) (*ExchangeRate, error)
	Update(ctx context.Context, exchangeRate *ExchangeRate) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, exchangeRates []*ExchangeRate) error
	UpdateBulk(ctx context.Context, exchangeRates []*ExchangeRate) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
