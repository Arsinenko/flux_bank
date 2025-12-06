package exchange_rate

import "time"

type ExchangeRate struct {
	RateID         int32
	BaseCurrency   *string
	TargetCurrency *string
	Rate           *string
	UpdatedAt      *time.Time
}
