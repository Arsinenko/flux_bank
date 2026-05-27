package actions

import (
	"orch-go/internal/domain/deposit"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
	"time"

	"github.com/shopspring/decimal"
)

type MakeDepositAction struct{}

func (a MakeDepositAction) Execute(ctx *simulation_context.SimpleSimulationContext, agent *agents.Agent) error {
	account, err := ctx.Services().AccountService.GetAccountById(ctx, agent.AccountId)
	if err != nil {
		return err
	}
	if account.Balance.LessThan(agent.Salary.Mul(decimal.NewFromFloat(0.9))) {
		return nil
	}
	depAmount := account.Balance.Mul(decimal.NewFromFloat(0.3))
	dep := deposit.Deposit{
		DepositID:    0,
		CustomerID:   agent.CustomerId,
		Amount:       depAmount.String(),
		InterestRate: "0.1",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(time.Hour * 24 * 365),
		Status:       "active",
	}
	_, err = ctx.Services().DepositService.CreateDeposit(ctx, &dep)
	if err != nil {
		return err
	}
	account.Balance = account.Balance.Sub(depAmount)
	err = ctx.Services().AccountService.UpdateAccount(ctx, account)
	if err != nil {
		return err
	}
	return nil
}
