package actions

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
)

type GetMoneyAction struct {
}

func (a GetMoneyAction) Execute(ctx *simulation_context.SimpleSimulationContext, agent *agents.Agent) error {
	account, err := ctx.Services().AccountService.GetAccountById(ctx, agent.AccountId)
	if err != nil {
		return err
	}
	account.Balance = account.Balance.Add(account.Balance)
	return ctx.Services().AccountService.UpdateAccount(ctx, account)
}
