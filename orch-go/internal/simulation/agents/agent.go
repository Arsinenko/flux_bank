package agents

import (
	"github.com/shopspring/decimal"
)

type Agent struct {
	CustomerId int32           `json:"customer_id"`
	AccountId  int32           `json:"account_id"`
	Salary     decimal.Decimal `json:"salary"`
}

func NewAgent(customerId int32, accountId int32) *Agent {
	return &Agent{
		CustomerId: customerId,
		AccountId:  accountId,
	}
}
