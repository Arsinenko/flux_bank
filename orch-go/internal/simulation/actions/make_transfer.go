package actions

import (
	"math/rand/v2"
	"orch-go/internal/domain/transaction"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
	"time"

	"github.com/shopspring/decimal"
)

type MakeTransfer struct {
}

func (a MakeTransfer) Execute(ctx *simulation_context.SimpleSimulationContext, agent *agents.Agent) error {
	ac, err := ctx.Services().AccountService.GetAccountById(ctx, agent.AccountId)
	if err != nil {
		return err
	}
	if ac.Balance.LessThan(agent.Salary.Mul(decimal.NewFromFloat(0.1))) {
		return nil
	}
	agents := ctx.Agents()
	targetAgent := agents[rand.IntN(len(agents))]

	min, max := 0.1, 0.3
	transferAmount := ac.Balance.Mul(decimal.NewFromFloat(min + rand.Float64()*(max-min)))
	now := time.Now()
	transfer := transaction.Transaction{
		TransactionID: 0,
		SourceAccount: &agent.AccountId,
		TargetAccount: &targetAgent.AccountId,
		Amount:        transferAmount,
		Currency:      "rub",
		Status:        &[]string{"active"}[0],
		CreatedAt:     &now,
	}
	_, err = ctx.Services().TransactionService.CreateTransaction(ctx, &transfer)
	if err != nil {
		return err
	}
	return nil

}
