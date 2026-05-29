package actions

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"

	"github.com/shopspring/decimal"
)

type GetCurrentMoneyAction struct{}

func (g GetCurrentMoneyAction) Execute(ctx *simulation_context.SimpleSimulationContext, agent agents.Agent) (decimal.Decimal, error) {
	account, err := ctx.Services().AccountService.GetAccountById(ctx, agent.AccountId)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return account.Balance, nil
}
